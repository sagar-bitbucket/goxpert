package models

import "time"

//UserAnswers Model
type UserAnswers struct {
	ID         int        `gorm:"id" json:"id"`
	UUID       string     `gorm:"uuid" json:"uuid"`
	UserID     int        `gorm:"user_id" json:"userID"`
	QuestionID int        `gorm:"question_id" json:"questionID"`
	Output     string     `gorm:"output" json:"output"`
	Answer     string     `gorm:"answer" json:"answer"`
	StartTime  time.Time  `gorm:"start_time" json:"startTime"`
	EndTime    *time.Time `gorm:"end_time" json:"endTime"`
	CreatedAt  time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt  *time.Time `gorm:"deleted_at" json:"-"`
}
