package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	log "gitlab.com/scalent/goxpert/logs"
	"gitlab.com/scalent/goxpert/models"
	repository "gitlab.com/scalent/goxpert/services/executor/pkg/v1/repositories"
)

// ExecutorService describes the service.
type ExecutorService interface {
	ExecuteProgram(context.Context, models.Program) (*[]models.ProgramOutput, error)
}

//ExecutorServiceImpl **
type ExecutorServiceImpl struct {
	executorRepo repository.ExecutorRepository
}

//NewExecutorServiceImpl inject depedancies user repositiory
func NewExecutorServiceImpl(executorRepo repository.ExecutorRepository) ExecutorService {
	return &ExecutorServiceImpl{executorRepo: executorRepo}
}

//ExecuteProgram **
func (b *ExecutorServiceImpl) ExecuteProgram(ctx context.Context, program models.Program) (*[]models.ProgramOutput, error) {
	response := []models.ProgramOutput{}
	structTest := models.ProgramOutput{}
	respStruct := models.TestOutput{}
	log.Logger(ctx).Info("ExecuteProgram ", program)
	// resp, isFailed, err := b.executorRepo.ExecuteProgram(ctx, program)
	resp, _, err := b.executorRepo.ExecuteProgram(ctx, program)

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(strings.Replace(resp.TestOutput, "\r\n", "\n", -1), "\n")
	fmt.Println(cmd)
	for _, arrResp := range cmd {
		json.Unmarshal([]byte(arrResp), &respStruct)

		if (respStruct.Action == "pass") || (respStruct.Action == "fail") && strings.Contains(respStruct.Test, "/") {

			strTest := strings.Replace(respStruct.Output, "\n", "", -1)
			strTest = strings.Trim(strTest, " ")
			sliceTest := strings.Split(strTest, " ")
			strTest = strings.Replace(strTest, sliceTest[0], "", -1)
			strTest = strings.Trim(strTest, " ")

			sliceTest = strings.Split(sliceTest[1], ":")

			structTest.Function = sliceTest[0]
			structTest.ProgramOutput = sliceTest[1]
			structTest.ExpectedOutput = sliceTest[2]

			structTest.TestResult = respStruct.Action
			structTest.TotalTime = respStruct.Elapsed

			response = append(response, structTest)

		}
	}

	return &response, nil
}
