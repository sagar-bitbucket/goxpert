package repositories

import (
	"github.com/jinzhu/gorm"
)

//QuestionsRepository implimets all methods in QuestionsRepository
type QuestionsRepository interface {
}

//QuestionsRepositoryImpl **
type QuestionsRepositoryImpl struct {
	dbConn *gorm.DB
}

//NewQuestionsRepositoryImpl inject dependancies of DataStore
func NewQuestionsRepositoryImpl(dbConn *gorm.DB) QuestionsRepository {
	return &QuestionsRepositoryImpl{dbConn: dbConn}
}
