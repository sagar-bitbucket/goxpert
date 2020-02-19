package endpoints

import (
	"github.com/gorilla/mux"
	handler "gitlab.com/scalent/goxpert/services/course/pkg/v1/handlers"
)

//NewQuestionsRoute All Application Routes Are defiend Here
func NewAnswersRoute(router *mux.Router, handler *handler.AnswersHandlersImpl) {
	router.HandleFunc("/v1/question/{question_id}/answer", handler.SubmitAnswer).Methods("POST")
}
