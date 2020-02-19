package services

import (
	repository "gitlab.com/scalent/goxpert/services/course/pkg/v1/repositories"
)

// QuestionsService describes the service.
type QuestionsService interface {
}

//QuestionsServiceImpl **
type QuestionsServiceImpl struct {
	questionsRepo repository.QuestionsRepository
}

//NewQuestionsServiceImpl inject depedancies user repositiory
func NewQuestionsServiceImpl(questionsRepo repository.QuestionsRepository) QuestionsService {
	return &QuestionsServiceImpl{questionsRepo: questionsRepo}
}
