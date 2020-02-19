package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.com/scalent/goxpert/models"
	"golang.org/x/crypto/bcrypt"
)

//UserRepository implimets all methods in UserRepository
type UserRepository interface {
	CreateUser(context.Context, models.User) (interface{}, error)
	Login(context.Context, models.UserLoginRequest) (models.User, error)
	GetUsers(context.Context, int, int) (models.GetUsersResponse, error)
	GetUserByID(context.Context, uint32) (models.User, error)
	UpdateUser(context.Context, uint32, models.UpdateUserRequest) error
	GetUserByEmail(context.Context, string) (models.User, error)
	DeleteUser(context.Context, uint32) error
	UpdatePassword(context.Context, uint32, string) error
	GetUserIdFromUUID(context.Context, string) (uint32, error)
	CreateResetPasswordToken(context.Context, string, string) error
	ValidateResetToken(context.Context, string) error
	GetEmailFromToken(context.Context, string) (models.ResetPasswordRequest, error)
}

//UserRepositoryImpl **
type UserRepositoryImpl struct {
	dbConn *gorm.DB
}

//NewUserRepositoryImpl inject dependancies of DataStore
func NewUserRepositoryImpl(dbConn *gorm.DB) UserRepository {
	return &UserRepositoryImpl{dbConn: dbConn}
}

//CreateUser create users entry in database
func (userRepositoryImpl UserRepositoryImpl) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	dbConn := userRepositoryImpl.dbConn
	var count int
	dbConn.Table("users").Where("email = ?", user.Email).Count(&count)

	if count != 0 {
		return nil, errors.New("Email address already in use")
	}

	if err := dbConn.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//Login returns jwt token
func (userRepositoryImpl UserRepositoryImpl) Login(ctx context.Context, user models.UserLoginRequest) (userData models.User, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Where("email=?", user.Email).First(&userData).Error
	if err != nil {
		return userData, errors.New("Invalid login details")
	}
	return userData, nil

}

func (userRepositoryImpl UserRepositoryImpl) GetUsers(ctx context.Context, page, limit int) (resp models.GetUsersResponse, err error) {
	offset := (page - 1) * limit

	dbConn := userRepositoryImpl.dbConn

	err = dbConn.Table("users").Where("user_type = ?", "user").Where("deleted_at IS NULL").Limit(limit).Offset(offset).Find(&resp.Users).Error
	dbConn.Table("users").Where("user_type = ?", "user").Where("deleted_at IS NULL").Count(&resp.Count)

	if err := userRepositoryImpl.dbConn.DB().Ping(); err != nil {
		return resp, err
	}
	return resp, nil
}

func (userRepositoryImpl UserRepositoryImpl) GetUserByID(ctx context.Context, id uint32) (resp models.User, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("users").Where("id=?", id).First(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (userRepositoryImpl UserRepositoryImpl) UpdateUser(ctx context.Context, id uint32, user models.UpdateUserRequest) error {
	dbConn := userRepositoryImpl.dbConn
	if user.Password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		user.Password = string(pass)
	}
	err := dbConn.Table("users").Where("id=?", id).Update(&user).Error
	return err
}

func (userRepositoryImpl UserRepositoryImpl) DeleteUser(ctx context.Context, id uint32) error {
	dbConn := userRepositoryImpl.dbConn
	err := dbConn.Table("users").Where("id=?", id).Update("deleted_at", time.Now()).Error
	return err
}

func (userRepositoryImpl UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (user models.User, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Where("email=?", email).First(&user).Error
	return user, err
}

func (userRepositoryImpl UserRepositoryImpl) UpdatePassword(ctx context.Context, id uint32, password string) error {
	dbConn := userRepositoryImpl.dbConn
	err := dbConn.Table("users").Where("id=?", id).Update("password", password).Error
	return err
}

func (userRepositoryImpl UserRepositoryImpl) GetUserIdFromUUID(ctx context.Context, uuid string) (id uint32, err error) {
	dbConn := userRepositoryImpl.dbConn
	user := models.User{}
	err = dbConn.Table("users").Select("id").Where("uuid=?", uuid).First(&user).Error

	if err != nil {
		return id, err
	}
	return user.ID, nil
}

func (userRepositoryImpl UserRepositoryImpl) CreateResetPasswordToken(ctx context.Context, email, token string) (err error) {
	dbConn := userRepositoryImpl.dbConn
	request := models.ResetPasswordRequest{}
	request.Email = email
	request.Token = token

	err = dbConn.Table("user_password_resets").Create(&request).Error

	if err != nil {

		request.CreatedAt = time.Now()

		err = dbConn.Table("user_password_resets").Where("email=?", email).Update(&request).Error
	}
	return err
}

func (userRepositoryImpl UserRepositoryImpl) ValidateResetToken(ctx context.Context, token string) (err error) {
	var count int
	now := time.Now().Add(-24 * time.Hour)
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("user_password_resets").Where("token = ?", token).Where("created_at > ?", now).Count(&count).Error

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Your reset password link is expired")
	}
	return nil
}

func (userRepositoryImpl UserRepositoryImpl) GetEmailFromToken(ctx context.Context, token string) (details models.ResetPasswordRequest, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("user_password_resets").Where("token = ?", token).First(&details).Error

	if err != nil {
		return details, err
	}
	return details, nil
}
