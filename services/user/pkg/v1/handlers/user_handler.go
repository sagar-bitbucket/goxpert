package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	log "gitlab.com/scalent/goxpert/logs"

	"net/http"

	"gitlab.com/scalent/goxpert/models"
	service "gitlab.com/scalent/goxpert/services/user/pkg/v1/services"
)

//UserHandlersImpl for handler Functions
type UserHandlersImpl struct {
	userSvc service.UsersService
}

//NewUserHandlerImpl inits dependancies for graphQL and Handlers
func NewUserHandlerImpl(userService service.UsersService) *UserHandlersImpl {
	return &UserHandlersImpl{userSvc: userService}
}

var httpErr models.HTTPErr

// Login Handler
/*
Description:- Login function used for getting jwt token.
Used By:- User and Admin
*/
func (userHandlersImpl UserHandlersImpl) Login(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	user := models.UserLoginRequest{}
	log.Logger(ctx).Info("in request")

	err := ReadInput(req.Body, &user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	resp, err := userHandlersImpl.userSvc.Login(ctx, user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}
	WriteOKResponse(w, resp)
}

// Logout for return token
func (userHandlersImpl UserHandlersImpl) Logout(w http.ResponseWriter, req *http.Request) {

	resp := ""
	WriteOKResponse(w, resp)

}

// GetUsers handler Function
func (userHandlersImpl UserHandlersImpl) GetUsers(w http.ResponseWriter, req *http.Request) {
	var strpage, strlimit string

	keys := req.URL.Query()
	pages, ok := keys["page"]
	if ok {
		strpage = pages[0]
	} else {
		strpage = "0"
	}
	limits, ok := keys["limit"]
	if ok {
		strlimit = limits[0]
	} else {
		strlimit = "10"
	}

	page, err := strconv.Atoi(strpage)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "page query parameter should be an int")
		return
	}

	limit, err := strconv.Atoi(strlimit)

	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, "page limit parameter should be an int")
		return
	}

	ctx := req.Context()
	resp, err := userHandlersImpl.userSvc.GetUsers(ctx, page, limit)

	if err != nil {

		WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(resp.Users) <= 0 {
		WriteHTTPError(w, http.StatusNotFound)
		return
	}
	WriteOKResponse(w, resp)
}

// GetUserByID handler Function
func (userHandlersImpl UserHandlersImpl) GetUserByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	resp, err := userHandlersImpl.userSvc.GetUserByID(ctx, id)

	if err != nil {
		WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}
	WriteOKResponse(w, resp)
}

// // GetUserByID handler Function
// func (userHandlersImpl UserHandlersImpl) GetUserByUUID(w http.ResponseWriter, req *http.Request) {
// 	ctx := req.Context()
// 	vars := mux.Vars(req)
// 	uuid := vars["uuid"]
// 	resp, err := userHandlersImpl.userSvc.GetUserByID(ctx, uuid)

// 	if err != nil {
// 		WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	WriteOKResponse(w, resp)
// }

// UpdateUser handler Function
func (userHandlersImpl UserHandlersImpl) UpdateUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	user := models.UpdateUserRequest{}
	log.Logger(ctx).Info("in request")

	err := ReadInput(req.Body, &user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	err = userHandlersImpl.userSvc.UpdateUser(ctx, id, user)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteOKResponse(w, "user updated successfully")
}

// CreateUser handler Function
func (userHandlersImpl UserHandlersImpl) CreateUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := models.User{}
	log.Logger(ctx).Info("in request")

	err := ReadInput(req.Body, &user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	resp, err := userHandlersImpl.userSvc.CreateUser(ctx, user)
	if err != nil {
		WriteCustomHTTPError(w, http.StatusConflict, err.Error())
		return
	}

	WriteOKResponse(w, resp)

}

// UserProfile handler Function
func (userHandlersImpl UserHandlersImpl) UserProfile(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	uuid, err := getUUIDFromToken(req)
	if err != nil {
		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// fmt.Println(claimsID)

	resp, err := userHandlersImpl.userSvc.GetUserByID(ctx, uuid)

	if err != nil {
		WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}
	WriteOKResponse(w, resp)
}

// UpdateProfile handler Function
func (userHandlersImpl UserHandlersImpl) UpdateProfile(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]

	user := models.UpdateUserRequest{}
	log.Logger(ctx).Info("in request", id)

	err := ReadInput(req.Body, &user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	err = userHandlersImpl.userSvc.UpdateUser(ctx, id, user)
	resp := ""

	WriteOKResponse(w, resp)
}

// DeleteUser handler Function
func (userHandlersImpl UserHandlersImpl) DeleteUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	log.Logger(ctx).Info("in request")

	err := userHandlersImpl.userSvc.DeleteUser(ctx, id)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	resp := ""
	WriteOKResponse(w, resp)

}

// ResendPassword handler Function Used for : Admin
func (userHandlersImpl UserHandlersImpl) ResendPassword(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	log.Logger(ctx).Info("in request")

	err := userHandlersImpl.userSvc.ResendPassword(ctx, id)
	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	resp := ""
	WriteOKResponse(w, resp)
}

// ForgotPassword handler Function
func (userHandlersImpl UserHandlersImpl) ForgotPassword(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := models.ForgotPasswordRequest{}
	log.Logger(ctx).Info("in request ForgotPassword")

	err := ReadInput(req.Body, &user)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	err = userHandlersImpl.userSvc.ForgotPassword(ctx, user)

	if err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
		return
	}

	resp := ""
	WriteOKResponse(w, resp)
}

// ResetPassword handler Function
func (userHandlersImpl UserHandlersImpl) ResetPassword(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	uniqueToken := vars["uniqueToken"]
	log.Logger(ctx).Info("in request")

	err := userHandlersImpl.userSvc.ResetPassword(ctx, uniqueToken)

	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := ""
	WriteOKResponse(w, resp)

}

// UpdatePassword handler Function
func (userHandlersImpl UserHandlersImpl) UpdatePassword(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	uniqueToken := vars["uniqueToken"]
	log.Logger(ctx).Info("in request")

	updateRequest := models.UpdatePasswordRequest{}

	err := ReadInput(req.Body, &updateRequest)
	if err != nil {
		WriteHTTPError(w, http.StatusBadRequest)
		return
	}

	err = userHandlersImpl.userSvc.UpdatePassword(ctx, uniqueToken, updateRequest)

	if err != nil {
		WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := ""
	WriteOKResponse(w, resp)

}

// GetUserCourseSections handler Function
func (userHandlersImpl UserHandlersImpl) GetUserCourseSections(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	//Get User Id From Tocken
	id, err := getUUIDFromToken(req)
	if err != nil {
		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//get CourseId from URL
	courseID, err := GetCourseID(req)
	if err != nil {
		CustomHTTPError(w, http.StatusBadRequest, "courseId in URL path should be integer")
		return
	}

	//get CourseId from URL
	sectionID, err := GetSectionID(req)
	if err != nil {
		CustomHTTPError(w, http.StatusBadRequest, "sectionId in URL path should be integer")
		return
	}
	request := models.GetUserCourseSectionsReq{}
	request.ID = id
	request.CourseID = courseID
	request.SectionID = sectionID
	resp, err := userHandlersImpl.userSvc.GetUserCourseSections(ctx, request)
	if err != nil {
		CustomHTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	endpointResp, err := json.Marshal(resp)
	if err != nil {
		HTTPError(w, http.StatusInternalServerError)
		return
	}

	w.Write(endpointResp)

}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}
