package endpoints

import (
	"github.com/gorilla/mux"
	handler "gitlab.com/scalent/goxpert/services/course/pkg/v1/handlers"
)

//NewQuestionsRoute All Application Routes Are defiend Here
func NewQuestionsRoute(router *mux.Router, handler *handler.QuestionsHandlersImpl) {
	router.HandleFunc("/v1/section/{section_id}/question", handler.CreateQuestion).Methods("POST")
	router.HandleFunc("/v1/section/{section_id}/section", handler.GetAllQuestion).Methods("GET")
	router.HandleFunc("/v1/section/{section_id}/section/{section_id}", handler.GetQuestion).Methods("GET")
	router.HandleFunc("/v1/section/{section_id}/section/{section_id}", handler.DeleteQuation).Methods("DELETE")
	router.HandleFunc("/v1/section/{section_id}/section/{section_id}", handler.UpdateQuestion).Methods("PATCH")
	router.HandleFunc("/v1/section/{section_id}/section", handler.UpdateQuestionSequence).Methods("PATCH")
}
