package handlers

import (
	"net/http"

	service "gitlab.com/scalent/goxpert/services/course/pkg/v1/services"
)

//AnswersHandlersImpl for handler Functions
type AnswersHandlersImpl struct {
	AnswersSvc service.AnswersService
}

//NewAnswersHandlerImpl inits dependancies for graphQL and Handlers
func NewAnswersHandlerImpl(answersService service.AnswersService) *AnswersHandlersImpl {
	return &AnswersHandlersImpl{AnswersSvc: answersService}
}

//var httpErr models.HTTPErr

//SubmitAnswer handler function
func (answersHandler AnswersHandlersImpl) SubmitAnswer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	question_id := req.FormValue("question_id")
	answer := req.FormValue("answer")
	uuid, err := getUUIDFromToken(req)

	//	fmt.Println(uuid)

	err = answersHandler.AnswersSvc.SubmitAnswer(ctx, question_id, answer, uuid)
	if err != nil {
		CustomHTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
