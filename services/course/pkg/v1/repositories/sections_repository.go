package repositories

import (
	"context"
	"time"

	log "gitlab.com/scalent/goxpert/logs"

	"github.com/jinzhu/gorm"
	"gitlab.com/scalent/goxpert/models"
)

const (
	//SectionRecordRetriveLimit retrive limits for no of records retrival  limit
	SectionRecordRetriveLimit = 10
)

//SectionsRepository implimets all methods in SectionsRepository
type SectionsRepository interface {
	CreateSection(context.Context, models.Sections) (interface{}, error)
	GetSections(context.Context, uint32, uint32) (*models.Sections, error)
	GetAllSections(context.Context, uint32) (*[]models.Sections, error)
	UpdateSections(context.Context, models.Sections) (interface{}, error)
	UpdateSectionsSequence(context.Context, models.ChnageSectionSequence) (interface{}, error)
	GetSectionsInfo(context.Context, models.SectionInfoReq) (interface{}, error)
}

//SectionsRepositoryImpl **
type SectionsRepositoryImpl struct {
	dbConn *gorm.DB
}

//NewSectionsRepositoryImpl inject dependancies of DataStore
func NewSectionsRepositoryImpl(dbConn *gorm.DB) SectionsRepository {
	return &SectionsRepositoryImpl{dbConn: dbConn}
}

//CreateSection Create section record in databse
func (secRepo SectionsRepositoryImpl) CreateSection(ctx context.Context,
	sectionReq models.Sections) (interface{}, error) {

	//databse connection
	conn := secRepo.dbConn
	createOn := time.Now().In(time.UTC)

	//record create Time
	sectionReq.CreatedAt = createOn

	//record update time change
	sectionReq.UpdatedAt = createOn

	d := conn.Create(&sectionReq)
	if d.Error != nil {
		log.Logger(ctx).Error(d.Error)
		return nil, d.Error
	}

	return sectionReq.ID, nil
}

//GetSections serch sections by id and return it
func (secRepo SectionsRepositoryImpl) GetSections(ctx context.Context,
	courseID uint32, sectionID uint32) (*models.Sections, error) {

	//databse connection
	conn := secRepo.dbConn
	section := models.Sections{}

	if err := conn.Where("id=?", sectionID).Find(&section).Error; err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &section, nil

}

//GetAllSections read all re
func (secRepo SectionsRepositoryImpl) GetAllSections(ctx context.Context,
	page uint32) (*[]models.Sections, error) {

	//databse connection
	conn := secRepo.dbConn
	sections := []models.Sections{}

	if err := conn.Limit(SectionRecordRetriveLimit).
		Offset(SectionRecordRetriveLimit * page).
		Find(&sections).Error; err != nil {

		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &sections, nil

}

//UpdateSections  dsjgf
func (secRepo SectionsRepositoryImpl) UpdateSections(ctx context.Context,
	section models.Sections) (interface{}, error) {

	//database connection pool
	conn := secRepo.dbConn

	if err := conn.Model(&models.Sections{}).Updates(&section).Error; err != nil {
		return nil, err
	}
	return section, nil
}

//UpdateSectionsSequence *
func (secRepo SectionsRepositoryImpl) UpdateSectionsSequence(ctx context.Context,
	req models.ChnageSectionSequence) (interface{}, error) {

	//database conne	qrwction pool
	conn := secRepo.dbConn

	conn.LogMode(true)

	txn := conn.Begin()
	if req.CurrentLocation > req.ReLocation {

		if err := txn.Table("sections").
			Where("sequence_number =? AND course_id = ?", req.CurrentLocation, req.CourseID).
			Update("sequence_number", req.ReLocation).
			Error; err != nil {
			log.Logger(ctx).Println(err)
			txn.Rollback()
			return nil, err
		}

		if err := txn.Table("sections").
			Where("sequence_number BETWEEN ? AND  ?  AND course_id = ? AND id !=?",
				req.ReLocation, req.CurrentLocation, req.CourseID, req.SectionID).
			Update("sequence_number", gorm.Expr("sequence_number + 1")).
			Error; err != nil {
			log.Logger(ctx).Println(err)
			txn.Rollback()
			return nil, err
		}
	} else {

		if err := txn.Table("sections").
			Where("sequence_number =? AND course_id = ?", req.CurrentLocation, req.CourseID).
			Update("sequence_number", req.ReLocation).
			Error; err != nil {
			log.Logger(ctx).Println(err)
			txn.Rollback()
			return nil, err
		}

		if err := txn.Table("sections").
			Where("sequence_number BETWEEN ? AND  ?  AND course_id = ? AND id !=?",
				req.CurrentLocation, req.ReLocation, req.CourseID, req.SectionID).
			Update("sequence_number", gorm.Expr("sequence_number - 1")).
			Error; err != nil {
			log.Logger(ctx).Println(err)
			txn.Rollback()
			return nil, err
		}

	}
	txn.Commit()
	return "Sections Sequence Updated", nil

}

//GetSectionsInfo *
func (secRepo SectionsRepositoryImpl) GetSectionsInfo(ctx context.Context,
	req models.SectionInfoReq) (interface{}, error) {

	//database connection pool
	conn := secRepo.dbConn

	secInfo := models.SectionInfo{}
	section := models.Sections{}

	type Value struct {
		SectionCompletionRer int `gorm:"section_completion_per"`
	}

	//SELECT ROUND((((SELECT COUNT(*) FROM questions as qu,user_answers as ans WHERE ans.question_id=qu.id AND ans.user_id=1 AND ans.is_correct=true)/(SELECT COUNT(*) FROM questions WHERE questions.section_id=1))*100)) as section_completion_per FROM questions as qu,user_answers as ans WHERE ans.question_id=qu.id AND ans.user_id=1

	txn := conn.Begin()
	query := `SELECT
		ROUND
		(
			(
				(
					(
						SELECT COUNT(*) 
						FROM questions as qu,user_answers as ans 
						WHERE ans.question_id=qu.id AND ans.user_id = ? 
						AND ans.is_correct=true
					)/
					(	SELECT COUNT(*) 
						FROM questions 
						WHERE questions.section_id = ?
					)
				)*100
			)
		)as section_completion_per 
		FROM questions as qu,user_answers as ans 
		WHERE ans.question_id=qu.id 
		AND ans.user_id=?`

	rows, err := txn.Debug().Raw(query, req.UserID, req.SectionID, req.UserID).Rows()

	if err != nil {
		log.Logger(ctx).Error(err)
		txn.Rollback()
		return nil, err

	}
	var count string
	for rows.Next() {
		rows.Scan(&count)

	}

	if err := txn.Where("id=?", req.SectionID).Find(&section).Error; err != nil {

		log.Logger(ctx).Error(err)
		txn.Rollback()
		return nil, err
	}
	txn.Commit()
	secInfo.Section = section
	secInfo.PercentageCompleted = count

	return secInfo, nil

}

//SELECT (SELECT  COUNT(*) FROM questions WHERE questions.section_id=1)
//SELECT  ((SELECT COUNT(*) FROM questions as qu,user_answers  as ans WHERE ans.question_id=qu.id AND ans.user_id=1)/(SELECT COUNT(*) FROM questions WHERE questions.section_id=1))*100)

//SELECT COUNT(*) FROM questions as qu,user_answers  as ans WHERE ans.question_id=qu.id AND ans.user_id=1
