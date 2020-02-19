package models

import (
	"time"
)

//Question Model
type Question struct {
	ID               int               `gorm:"id" json:"id"`
	SectionID        int               `gorm:"section_id" json:"sectionID"`
	SequenceNumber   int               `gorm:"sequence_number" json:"sequenceNumber"`
	QuestionType     string            `gorm:"question_type" json:"questionType"`
	QuestionTitle    string            `gorm:"question_title" json:"questionTitle"`
	ProblemStatement string            `gorm:"problemStatement" json:"problemStatement"`
	TestCases        string            `gorm:"testCases" json:"testCases"`
	CreatedAt        time.Time         `gorm:"created_at" json:"-"`
	UpdatedAt        time.Time         `gorm:"updated_at" json:"-"`
	DeletedAt        *time.Time        `gorm:"deleted_at" json:"-"`
	Options          []QuestionOptions `gorm:"-" json:"options"`
}

//TestCases **
type TestCases struct {
	TestCaseID   uint
	TestCaseName string
}

//QuestionOptions **
type QuestionOptions struct {
	ID         uint
	QuestionID int
	Name       string
	Containt   string
	IsAnswer   string
}

//QuestionTree Struct
type QuestionTree struct {
	ID             int    `json:"id"`
	SequenceNumber int    `json:"sequenceNumber"`
	QuestionType   string `json:"questionType"`
	QuestionTitle  string `json:"questionTitle"`
}
