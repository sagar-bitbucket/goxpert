package endpoints

import (
	"github.com/gorilla/mux"
	handler "gitlab.com/scalent/goxpert/services/user/pkg/v1/handlers"
)

//NewUserRoute All Application Routes Are defiend Here
func NewUserRoute(router *mux.Router, handler *handler.UserHandlersImpl) {
	//** login user - Common //
	router.HandleFunc("/login", handler.Login).Methods("POST")
	//** logout user - Common //
	router.HandleFunc("/logout", handler.Logout).Methods("GET")
	//** Get all users-For Admin //
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	//** Get user by id- for Admin//
	router.HandleFunc("/user/{id}", handler.GetUserByID).Methods("GET")
	//** Get user by id- for Internal call//
	//router.HandleFunc("/user/details/{uuid}", handler.GetUserByUUID).Methods("GET")
	//** Get User By Token - for user//
	router.HandleFunc("/user", handler.UserProfile).Methods("GET")
	//** Update User by token - for User//
	router.HandleFunc("/user", handler.UpdateProfile).Methods("PUT")
	//** Update User - for Admin//
	router.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PATCH")
	//** Get User By Token - for User//
	router.HandleFunc("/user", handler.CreateUser).Methods("POST")
	//** Delete User - for Admin//
	router.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")
	//** Resend login details - for Admin//
	router.HandleFunc("/user/resendpass/{id}", handler.ResendPassword).Methods("PATCH")
	//** Forgot Password - for User//
	router.HandleFunc("/user/forgotpass", handler.ForgotPassword).Methods("POST")
	//** Reset password - for User//
	router.HandleFunc("/user/resetpass/{uniqueToken}", handler.ResetPassword).Methods("GET")
	//** Update password with confirmation - for User//
	router.HandleFunc("/user/resetpass/{uniqueToken}", handler.UpdatePassword).Methods("POST")
	//** Update password with confirmation - for User//
	router.HandleFunc("/v1/user/course/{courseId}/sections/{sectionId}", handler.GetUserCourseSections).Methods("GET")

}
