package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	log "gitlab.com/scalent/goxpert/logs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	//CourseID store courseID
	CourseID = "courseId"

	//SectionID stores sectionID
	SectionID = "sectionId"

	//JWTSecretKey is used to create tocken
	JWTSecretKey = `95a31f74-a1cd-4321-8a9d-bdb0735e445a`
)

var statusText = map[int]string{
	http.StatusContinue:           "Continue",
	http.StatusSwitchingProtocols: "Switching Protocols",
	http.StatusProcessing:         "Processing",
	//	http.StatusEarlyHints:           "Early Hints",

	http.StatusOK:                   "OK",
	http.StatusCreated:              "Created",
	http.StatusAccepted:             "Accepted",
	http.StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	http.StatusNoContent:            "No Content",
	http.StatusResetContent:         "Reset Content",
	http.StatusPartialContent:       "Partial Content",
	http.StatusMultiStatus:          "Multi-Status",
	http.StatusAlreadyReported:      "Already Reported",
	http.StatusIMUsed:               "IM Used",

	http.StatusMultipleChoices:   "Multiple Choices",
	http.StatusMovedPermanently:  "Moved Permanently",
	http.StatusFound:             "Found",
	http.StatusSeeOther:          "See Other",
	http.StatusNotModified:       "Not Modified",
	http.StatusUseProxy:          "Use Proxy",
	http.StatusTemporaryRedirect: "Temporary Redirect",
	http.StatusPermanentRedirect: "Permanent Redirect",

	http.StatusBadRequest:                    "Bad Request",
	http.StatusUnauthorized:                  "Unauthorized",
	http.StatusPaymentRequired:               "Payment Required",
	http.StatusForbidden:                     "Forbidden",
	http.StatusNotFound:                      "Not Found",
	http.StatusMethodNotAllowed:              "Method Not Allowed",
	http.StatusNotAcceptable:                 "Not Acceptable",
	http.StatusProxyAuthRequired:             "Proxy Authentication Required",
	http.StatusRequestTimeout:                "Request Timeout",
	http.StatusConflict:                      "Conflict",
	http.StatusGone:                          "Gone",
	http.StatusLengthRequired:                "Length Required",
	http.StatusPreconditionFailed:            "Precondition Failed",
	http.StatusRequestEntityTooLarge:         "Request Entity Too Large",
	http.StatusRequestURITooLong:             "Request URI Too Long",
	http.StatusUnsupportedMediaType:          "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	http.StatusExpectationFailed:             "Expectation Failed",
	http.StatusTeapot:                        "I'm a teapot",
	http.StatusMisdirectedRequest:            "Misdirected Request",
	http.StatusUnprocessableEntity:           "Unprocessable Entity",
	http.StatusLocked:                        "Locked",
	http.StatusFailedDependency:              "Failed Dependency",
	http.StatusTooEarly:                      "Too Early",
	http.StatusUpgradeRequired:               "Upgrade Required",
	http.StatusPreconditionRequired:          "Precondition Required",
	http.StatusTooManyRequests:               "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	http.StatusInternalServerError:           "Internal Server Error",
	http.StatusNotImplemented:                "Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

//HTTPError return HTTP Error Message
//@TODO:Remove this func, other modified code added bellow
func HTTPError(w http.ResponseWriter, statusCode int) {
	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg, statusCode)
}

//CustomHTTPError **
//@TODO:Remove this func, other modified code added bellow
func CustomHTTPError(w http.ResponseWriter, statusCode int, CustomErrMessage string) {

	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg+" : "+CustomErrMessage, statusCode)
}

//GetCourseID get CourseID and returns its value in int format
func GetCourseID(req *http.Request) (uint32, error) {

	params := mux.Vars(req)
	id := params[CourseID]
	fmt.Println(id)
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Logger(req.Context()).Error("Error in courseId Parsing", err)
		return 0, errors.New("courseId in URL path should be integer")
	}
	return uint32(uintID), nil
}

//GetSectionID get CourseID and returns its value in int format
func GetSectionID(req *http.Request) (uint32, error) {

	params := mux.Vars(req)
	id := params[SectionID]
	fmt.Println(id)
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Logger(req.Context()).Error("Error in sectionId Parsing", err)
		return 0, errors.New("sectionId in URL path should be integer")
	}
	return uint32(uintID), nil
}

func getUUIDFromToken(req *http.Request) (string, error) {
	var uuid string
	ctx := req.Context()
	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, "Bearer")
		reqToken := bearerToken[1]
		reqToken = strings.TrimSpace(reqToken)
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTSecretKey), nil
		})
		if err != nil {
			log.Logger(ctx).Error("Error in sectionId Parsing", err)
			return "", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Logger(ctx).Error("Error in sectionId Parsing", err)
			return "", errors.New("erros  in type assertio of jwt.MapClaims ")
		}
		//here uuid is just an key to store user Id Into table
		uuid := claims["uuid"].(string)
		fmt.Println(claims)
		return uuid, nil
	}
	return uuid, nil
}

//ReadInput from the body
func ReadInput(rBody io.ReadCloser, input interface{}) error {
	decoder := json.NewDecoder(rBody)
	err := decoder.Decode(input)
	return err
}

//WriteOKResponse as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		WriteHTTPError(w, http.StatusInternalServerError)
	}
}

//WriteHTTPError return HTTP Error Message
func WriteHTTPError(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg, statusCode)
}

//WriteCustomHTTPError **
func WriteCustomHTTPError(w http.ResponseWriter, statusCode int, CustomErrMessage string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	errorMsg := statusText[statusCode]
	http.Error(w, errorMsg+" : "+CustomErrMessage, statusCode)
}
