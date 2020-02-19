package models

import "time"

//UserResults Model
type UserResults struct {
	ID            int       `gorm:"id" json:"id"`
	UUID          string    `gorm:"uuid" json:"uuid"`
	UserID        int       `gorm:"user_id" json:"userID"`
	UserAnswersID int       `gorm:"user_answers_id" json:"answerID"`
	IsCorrect     bool      `gorm:"is_correct" json:"isCorrect"`
	Coverage      float64   `gorm:"coverage" json:"coverage"`
	CreatedAt     time.Time `gorm:"created_at" json:"createdAt"`
}
