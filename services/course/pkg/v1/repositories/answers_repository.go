package repositories

import (
	"context"

	"github.com/jinzhu/gorm"
	"gitlab.com/scalent/goxpert/models"
)

//AnswersRepository implimets all methods in AnswersRepository
type AnswersRepository interface {
	SubmitAnswer(context.Context, string, string, string, uint32) error
	GetQuetionByID(context.Context, string) (models.Question, error)
	GetCorrectOptionsByID(context.Context, string) (models.QuestionOptions, error)
}

//AnswersRepositoryImpl **
type AnswersRepositoryImpl struct {
	dbConn *gorm.DB
}

//NewAnswersRepositoryImpl inject dependancies of DataStore
func NewAnswersRepositoryImpl(dbConn *gorm.DB) AnswersRepository {
	return &AnswersRepositoryImpl{dbConn: dbConn}
}

func (answersRepository AnswersRepositoryImpl) SubmitAnswer(ctx context.Context, id, answer, is_correct string, user_id uint32) error {

	//conn := answersRepository.dbConn
	return nil
}

func (answersRepository AnswersRepositoryImpl) GetQuetionByID(ctx context.Context, id string) (resp models.Question, err error) {

	dbConn := answersRepository.dbConn
	err = dbConn.Table("questions").Where("id=?", id).First(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (answersRepository AnswersRepositoryImpl) GetCorrectOptionsByID(ctx context.Context, id string) (resp models.QuestionOptions, err error) {

	dbConn := answersRepository.dbConn
	err = dbConn.Table("question_options").Where("question_id=?", id).Where("is_correct=?", "yes").First(&resp).Error

	if err != nil {
		return resp, err
	}
	return resp, err
}
