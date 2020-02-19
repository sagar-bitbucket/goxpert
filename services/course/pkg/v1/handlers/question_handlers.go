package handlers

import (
	"net/http"

	"gitlab.com/scalent/goxpert/models"
	service "gitlab.com/scalent/goxpert/services/course/pkg/v1/services"
)

//QuestionsHandlersImpl for handler Functions
type QuestionsHandlersImpl struct {
	questionsSvc service.QuestionsService
}

//NewQuestionsHandlerImpl inits dependancies for graphQL and Handlers
func NewQuestionsHandlerImpl(questionsService service.QuestionsService) *QuestionsHandlersImpl {
	return &QuestionsHandlersImpl{questionsSvc: questionsService}
}

var httpErr models.HTTPErr

//CreateQuestion handler function
func (questionsHandlersImpl QuestionsHandlersImpl) CreateQuestion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//UpdateQuestion handler function
func (questionsHandlersImpl QuestionsHandlersImpl) UpdateQuestion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//GetQuestion handler function
func (questionsHandlersImpl QuestionsHandlersImpl) GetQuestion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//GetAllQuestion handler function
func (questionsHandlersImpl QuestionsHandlersImpl) GetAllQuestion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//DeleteQuation handler function
func (questionsHandlersImpl QuestionsHandlersImpl) DeleteQuation(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//UpdateQuestionSequence handler function
func (questionsHandlersImpl QuestionsHandlersImpl) UpdateQuestionSequence(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}
