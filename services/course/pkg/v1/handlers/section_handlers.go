package handlers

import (
	"net/http"
	"strconv"

	"gitlab.com/scalent/goxpert/models"

	service "gitlab.com/scalent/goxpert/services/course/pkg/v1/services"
)

//SectionsHandlersImpl for handler Functions
type SectionsHandlersImpl struct {
	sectionsSvc service.SectionsService
}

//NewSectionsHandlerImpl inits dependancies for graphQL and Handlers
func NewSectionsHandlerImpl(sectionsService service.SectionsService) *SectionsHandlersImpl {
	return &SectionsHandlersImpl{sectionsSvc: sectionsService}
}

//CreateSections handler function
func (sectionHaddler SectionsHandlersImpl) CreateSections(w http.ResponseWriter, req *http.Request) {

	section := models.Sections{}

	//requst Context
	ctx := req.Context()

	courseID, err := GetCourseID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	section.CourseID = courseID

	if err = ReadInput(req.Body, &section); err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	resp, err := sectionHaddler.sectionsSvc.CreateSection(ctx, section)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)

}

//UpdateSections handler function
func (sectionHaddler SectionsHandlersImpl) UpdateSections(w http.ResponseWriter, req *http.Request) {

	section := models.Sections{}

	//requst Context
	ctx := req.Context()

	courseID, err := GetCourseID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	sectionID, err := GetSectionID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "sectionId in URL path should be integer")
		return
	}

	section.CourseID = courseID
	section.ID = sectionID
	if err = ReadInput(req.Body, &section); err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	resp, err := sectionHaddler.sectionsSvc.UpdateSections(ctx, section)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)
}

//GetSections handler function
func (sectionHaddler SectionsHandlersImpl) GetSections(w http.ResponseWriter, req *http.Request) {
	//requst Context
	ctx := req.Context()

	courseID, err := GetCourseID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	sectionID, err := GetSectionID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "sectionId in URL path should be integer")
		return
	}

	resp, err := sectionHaddler.sectionsSvc.GetSections(ctx, courseID, sectionID)

	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)
}

//GetSectionsInfo handler function
func (sectionHaddler SectionsHandlersImpl) GetSectionsInfo(w http.ResponseWriter, req *http.Request) {
	//requst Context
	ctx := req.Context()

	userID, err := GetParamID(req, "user_id")
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}
	courseID, err := GetCourseID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	sectionID, err := GetSectionID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "sectionId in URL path should be integer")
		return
	}

	secReq := models.SectionInfoReq{}
	secReq.CourseID = courseID
	secReq.SectionID = sectionID
	secReq.UserID = userID

	resp, err := sectionHaddler.sectionsSvc.GetSectionsInfo(ctx, secReq)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)
}

//GetAllSections handler function
func (sectionHaddler SectionsHandlersImpl) GetAllSections(w http.ResponseWriter, req *http.Request) {
	//requst Context
	var strpage string
	keys := req.URL.Query()
	pages, ok := keys["page"]
	if ok {
		strpage = pages[0]
	} else {
		strpage = "0"
	}

	// String to integer Conversion
	pageInt, err := strconv.ParseInt(strpage, 10, 32)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "page query parameter should be an int")
		return
	}

	ctx := req.Context()
	resp, err := sectionHaddler.sectionsSvc.GetAllSections(ctx, uint32(pageInt))
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	//Record Not Found Error
	if len(*resp) <= 0 {
		WriteHTTPError(w, http.StatusNotFound)
		return
	}

	WriteOKResponse(w, resp)
}

//UpdateSectionsSequence handler function
func (sectionHaddler SectionsHandlersImpl) UpdateSectionsSequence(w http.ResponseWriter, req *http.Request) {
	//requst Context
	ctx := req.Context()

	courseID, err := GetCourseID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	sectionID, err := GetSectionID(req)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "sectionId in URL path should be integer")
		return
	}

	chSec := models.ChnageSectionSequence{}

	chSec.CourseID = courseID
	chSec.SectionID = sectionID
	if err = ReadInput(req.Body, &chSec); err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}
	resp, err := sectionHaddler.sectionsSvc.UpdateSectionsSequence(ctx, chSec)

	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	WriteOKResponse(w, resp)
}
