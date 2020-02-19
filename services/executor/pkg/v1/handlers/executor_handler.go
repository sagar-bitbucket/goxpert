package handlers

import (
	"encoding/json"

	log "gitlab.com/scalent/goxpert/logs"

	"net/http"

	"gitlab.com/scalent/goxpert/models"
	service "gitlab.com/scalent/goxpert/services/executor/pkg/v1/services"
)

//ExecutorHandlersImpl for handler Functions
type ExecutorHandlersImpl struct {
	executorSvc service.ExecutorService
}

//NewExecutorHandlerImpl inits dependancies for graphQL and Handlers
func NewExecutorHandlerImpl(executorService service.ExecutorService) *ExecutorHandlersImpl {
	return &ExecutorHandlersImpl{executorSvc: executorService}
}

var httpErr models.HTTPErr

// ExecuteProgram handler Function
func (executorHandlersImpl ExecutorHandlersImpl) ExecuteProgram(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	program := models.Program{}
	//	fmt.Println(req.FormValue("answer"))

	program.Code = req.FormValue("answer")
	program.TestCases = req.FormValue("testCases")

	log.Logger(ctx).Info("in request")
	//err := json.NewDecoder(req.Body).Decode(&program)

	//fmt.Println(program, "here")

	// if err != nil {

	// 	httpErr.Message = err.Error()
	// 	json.NewEncoder(w).Encode(httpErr)
	// 	writeResponse(w, http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := executorHandlersImpl.executorSvc.ExecuteProgram(ctx, program)
	if err != nil {
		httpErr.Message = err.Error()
		json.NewEncoder(w).Encode(httpErr)
		writeResponse(w, http.StatusInternalServerError)
		return
	}

	endpointResp, err := json.Marshal(resp)
	if err != nil {
		httpErr.Message = err.Error()
		json.NewEncoder(w).Encode(httpErr)
		writeResponse(w, http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
	w.Write(endpointResp)

}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}
