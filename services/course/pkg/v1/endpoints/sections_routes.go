package endpoints

import (
	"github.com/gorilla/mux"
	handler "gitlab.com/scalent/goxpert/services/course/pkg/v1/handlers"
)

//NewSectionsRoute All Application Routes Are defiend Here
func NewSectionsRoute(router *mux.Router, handler *handler.SectionsHandlersImpl) {

	router.HandleFunc("/v1/course/{courseId}/section", handler.CreateSections).Methods("POST")
	router.HandleFunc("/v1/course/{courseId}/section/{sectionId}", handler.UpdateSections).Methods("PATCH")
	router.HandleFunc("/v1/course/{courseId}/section/{sectionId}", handler.GetSections).Methods("GET")
	router.HandleFunc("/v1/course/{courseId}/section/{sectionId}/info", handler.GetSectionsInfo).Methods("GET")
	router.HandleFunc("/v1/course/{courseId}/section", handler.GetAllSections).Methods("GET")
	router.HandleFunc("/v1/course/{courseId}/section/{sectionId}/sequence", handler.UpdateSectionsSequence).Methods("PATCH")

}
