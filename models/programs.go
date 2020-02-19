package models

import "time"

//Program Model
type Program struct {
	Code      string `json:"code"`
	TestCases string `json:"testCases"`
}

//ProgramResponse Model
type ProgramResponse struct {
	ProgramOutput string `json:"programOutput"`
	TestOutput    string `json:"testOutput"`
}

type TestOutput struct {
	Time    time.Time `json:"-"`
	Action  string    `json:"action"`
	Package string    `json:"package"`
	Test    string    `json:"test"`
	Elapsed float64   `json:"elapsed"`
	Output  string    `json:"output"`
}

type ProgramOutput struct {
	TestName       string  `json:"test_name"`
	TestResult     string  `json:"test_result"`
	Function       string  `json:"function"`
	ProgramOutput  string  `json:"program_output"`
	ExpectedOutput string  `json:"expected_output"`
	TotalTime      float64 `json:"total_time"`
}
