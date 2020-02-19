package endpoints

import (
	"github.com/gorilla/mux"
	handler "gitlab.com/scalent/goxpert/services/executor/pkg/v1/handlers"
)

//NewExecutorRoute All Application Routes Are defiend Here
func NewExecutorRoute(router *mux.Router, handler *handler.ExecutorHandlersImpl) {
	router.HandleFunc("/program", handler.ExecuteProgram).Methods("POST")
}
