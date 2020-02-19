package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//User Model
type User struct {
	ID          uint32     `gorm:"id" json:"id"`
	Email       string     `gorm:"email" json:"email"`
	Name        string     `gorm:"name" json:"name"`
	UUID        string     `gorm:"uuid" json:"uuid"`
	Password    string     `gorm:"password" json:"-"`
	Designation string     `gorm:"designation" json:"designation"`
	EmpID       string     `gorm:"emp_id" json:"empID"`
	UserType    string     `gorm:"user_type" json:"userType"`
	UserStatus  string     `gorm:"user_status" json:"userStatus"`
	CreatedAt   time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"deleted_at" json:"-"`
}

//UserLoginRequest Model
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required" email:"true"`
	Password string `json:"password" validate:"required"`
}

//UserLoginResponse Model
type UserLoginResponse struct {
	Token    string `json:"token"`
	UserType string `json:"usertype"`
}

//GetUsersResponse Model
type GetUsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

//UpdateUserRequest Model
type UpdateUserRequest struct {
	Name        string `gorm:"name" json:"name,omitempty"`
	Password    string `gorm:"password" json:"password,omitempty"`
	Email       string `gorm:"email" json:"email,omitempty"`
	Designation string `gorm:"designation" json:"designation,omitempty"`
	EmpID       string `gorm:"emp_id" json:"empID,omitempty"`
}

type ForgotPasswordRequest struct {
	Email string `gorm:"email" json:"email"`
}

// JWTClaims describes Claims in token.
type JWTClaims struct {
	UUID     string `json:"uuid"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

type ResetPasswordRequest struct {
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
}

type UpdatePasswordRequest struct {
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

//GetUserCourseSectionsReq **
type GetUserCourseSectionsReq struct {
	ID        string
	CourseID  uint32
	SectionID uint32
}
