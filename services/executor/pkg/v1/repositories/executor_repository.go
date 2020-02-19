package repository

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"gitlab.com/scalent/goxpert/models"
)

//ExecutorRepository implimets all methods in ExecutorRepository
type ExecutorRepository interface {
	ExecuteProgram(context.Context, models.Program) (*models.ProgramResponse, bool, error)
}

//ExecutorRepositoryImpl **
type ExecutorRepositoryImpl struct {
	// dbConn *gorm.DB
}

//NewExecutorRepositoryImpl inject dependancies of DataStore
func NewExecutorRepositoryImpl() ExecutorRepository {
	return &ExecutorRepositoryImpl{}
}

func (executorRepositoryImpl ExecutorRepositoryImpl) ExecuteProgram(ctx context.Context, program models.Program) (*models.ProgramResponse, bool, error) {
	var isError bool
	output := models.ProgramResponse{}
	tempDirectory, err := ioutil.TempDir("programs", "programs")

	if err != nil {
		return nil, isError, errors.New("Unable to create temp directory")
	}

	defer os.RemoveAll(tempDirectory)

	progFilePath := tempDirectory + "/" + "main.go"
	testFilePath := tempDirectory + "/" + "main_test.go"
	progFilePtr, err := os.OpenFile(progFilePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, isError, errors.New("Unable to open main file")
	}

	testFilePtr, err := os.OpenFile(testFilePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, isError, errors.New("Unable to open test file")
	}

	defer progFilePtr.Close()
	defer testFilePtr.Close()

	ioutil.WriteFile(progFilePath, []byte(program.Code), 0777)
	ioutil.WriteFile(testFilePath, []byte(program.TestCases), 0777)
	fmt.Println("before exectution")

	cmd, _ := exec.Command("go", "test", "-json", "./"+tempDirectory).CombinedOutput()

	if err != nil {
		isError = true
	}

	testOutput := string(cmd)

	cmd, err = exec.Command("go", "run", "./"+tempDirectory).CombinedOutput()

	if err != nil {
		return nil, isError, err
	}

	programOutput := string(cmd)

	output.ProgramOutput = programOutput
	output.TestOutput = testOutput

	return &output, isError, nil
}
