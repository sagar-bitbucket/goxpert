package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/scalent/goxpert/email"
	log "gitlab.com/scalent/goxpert/logs"
	"golang.org/x/crypto/bcrypt"

	uuid "github.com/google/uuid"
	repository "gitlab.com/scalent/goxpert/services/user/pkg/v1/repositories"

	sendmail "gitlab.com/scalent/goxpert/email"
	"gitlab.com/scalent/goxpert/models"
)

// Claims describes Claims in token.
type Claims struct {
	ID       uint32 `json:"id"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

// UsersService describes the service.
type UsersService interface {
	CreateUser(ctx context.Context, createReq models.User) (interface{}, error)
	Login(context.Context, models.UserLoginRequest) (models.UserLoginResponse, error)
	GetUsers(context.Context, int, int) (models.GetUsersResponse, error)
	GetUserByID(context.Context, string) (models.User, error)
	UpdateUser(context.Context, string, models.UpdateUserRequest) error
	DeleteUser(context.Context, string) error
	ResendPassword(context.Context, string) error
	ForgotPassword(context.Context, models.ForgotPasswordRequest) error
	ResetPassword(context.Context, string) error
	UpdatePassword(context.Context, string, models.UpdatePasswordRequest) error
	GetUserCourseSections(context.Context, models.GetUserCourseSectionsReq) (interface{}, error)
}

//UsersServiceImpl **
type UsersServiceImpl struct {
	//mysql database oprations dependancies
	userRepo repository.UserRepository

	//call to external section service repository
	sectionCallRepo repository.SectionServiceRepo

	//dependancie for send in blue service provider
	sendBlue *email.SendInBlue
}

//NewUserServiceImpl inject depedancies user repositiory
func NewUserServiceImpl(userRepo repository.UserRepository,
	sendInBlue *email.SendInBlue,
	sectionServiceRepo repository.SectionServiceRepo) UsersService {

	return &UsersServiceImpl{
		userRepo:        userRepo,
		sendBlue:        sendInBlue,
		sectionCallRepo: sectionServiceRepo,
	}
}

//CreateUser **
func (b *UsersServiceImpl) CreateUser(ctx context.Context, user models.User) (interface{}, error) {

	log.Logger(ctx).Info("CreateUser ", user)
	_, err := b.userRepo.GetUserByEmail(ctx, user.Email)

	if err == nil {
		return user, errors.New("Email address already in use")
	}

	UUID := uuid.New()
	user.UUID = UUID.String()
	user.UserType = "user"
	user.UserStatus = "created"
	password := generateRandomString(8)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHash)

	resp, err := b.userRepo.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	// **send email service

	basepathStr := getServiceBasepath()
	TemplateData := make(map[string]string)
	TemplateData["email"] = user.Email
	TemplateData["password"] = password
	TemplateData["URL"] = basepathStr + "/login"

	TemplateName := "registration.html"

	data := sendmail.SendBlueData{}
	data.FromEmail = "info@goxpert.com"
	data.Subject = "Registration Successful"
	data.ToEmail = user.Email
	data.ToName = "Mr./Ms"
	data.Content = data.TemplateRender(ctx, TemplateData, TemplateName)

	b.sendBlue.SendMail(ctx, data)

	return resp, nil
}

//Login **
func (b *UsersServiceImpl) Login(ctx context.Context, user models.UserLoginRequest) (resp models.UserLoginResponse, err error) {
	log.Logger(ctx).Info("LoginUser ", user)
	userData, err := b.userRepo.Login(ctx, user)
	if err != nil {
		return resp, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return resp, errors.New("Invalid login details")
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.JWTClaims{
		UserType: userData.UserType,
		UUID:     userData.UUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenDetails := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenDetails.SignedString([]byte("95a31f74-a1cd-4321-8a9d-bdb0735e445a"))

	if err != nil {
		return resp, errors.New("Invalid login details")
	}

	resp.Token = token
	resp.UserType = userData.UserType
	return resp, nil
}

//GetUsers **
func (b *UsersServiceImpl) GetUsers(ctx context.Context, page, limit int) (resp models.GetUsersResponse, err error) {

	log.Logger(ctx).Info("GetUsers ")
	resp, err = b.userRepo.GetUsers(ctx, page, limit)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

//GetUserByID **
func (b *UsersServiceImpl) GetUserByID(ctx context.Context, uuid string) (resp models.User, err error) {
	log.Logger(ctx).Info(" GetUserByID")
	id, err := b.userRepo.GetUserIdFromUUID(ctx, uuid)

	if err != nil {
		return resp, err
	}

	resp, err = b.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return resp, err
	}

	return resp, err
}

//UpdateUser **
func (b *UsersServiceImpl) UpdateUser(ctx context.Context, uuid string, user models.UpdateUserRequest) error {
	id, err := b.userRepo.GetUserIdFromUUID(ctx, uuid)

	if err != nil {
		return err
	}

	err = b.userRepo.UpdateUser(ctx, id, user)
	return err
}

//DeleteUser **
func (b *UsersServiceImpl) DeleteUser(ctx context.Context, uuid string) error {
	id, err := b.userRepo.GetUserIdFromUUID(ctx, uuid)

	if err != nil {
		return err
	}

	err = b.userRepo.DeleteUser(ctx, id)
	return err
}

//ResendPassword **
func (b *UsersServiceImpl) ResendPassword(ctx context.Context, uuid string) error {

	id, err := b.userRepo.GetUserIdFromUUID(ctx, uuid)
	fmt.Println(err)

	if err != nil {
		return err
	}

	user, err := b.userRepo.GetUserByID(ctx, id)
	fmt.Println(id)
	if err != nil {
		return err
	}

	password := generateRandomString(8)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	err = b.userRepo.UpdatePassword(ctx, id, string(passwordHash))

	basepathStr := getServiceBasepath()
	TemplateData := make(map[string]string)
	TemplateData["Email"] = user.Email
	TemplateData["Password"] = password
	TemplateData["User"] = user.Name
	TemplateData["URL"] = basepathStr + "/login"

	TemplateName := "registration.html"

	data := sendmail.SendBlueData{}
	data.FromEmail = "info@goxpert.com"
	data.Subject = "Registration Successful"
	data.ToEmail = user.Email
	data.ToName = "Mr./Ms"
	data.Content = data.TemplateRender(ctx, TemplateData, TemplateName)

	b.sendBlue.SendMail(ctx, data)

	// **send email service

	return err
}

//ForgotPassword **
func (b *UsersServiceImpl) ForgotPassword(ctx context.Context,
	user models.ForgotPasswordRequest) error {

	log.Logger(ctx).Info("CreateUser ", user)
	_, err := b.userRepo.GetUserByEmail(ctx, user.Email)

	if err != nil {
		return err
	}

	token := generateRandomString(50)
	token = strings.ToLower(token)

	err = b.userRepo.CreateResetPasswordToken(ctx, user.Email, token)

	if err != nil {
		return err
	}

	basepathStr := getServiceBasepath()
	TemplateData := make(map[string]string)
	TemplateName := "forgot_password.html"

	TemplateData["URL"] = basepathStr + "/user/resetpass/" + token

	data := sendmail.SendBlueData{}
	data.FromEmail = "info@goxpert.com"
	data.Subject = "Forget Password"
	data.ToEmail = user.Email
	data.ToName = "Mr./Ms"
	data.Content = data.TemplateRender(ctx, TemplateData, TemplateName)

	b.sendBlue.SendMail(ctx, data)

	return nil
}

//ResetPassword **
func (b *UsersServiceImpl) ResetPassword(ctx context.Context, token string) error {

	log.Logger(ctx).Info("ResetPassword")
	err := b.userRepo.ValidateResetToken(ctx, token)
	return err
}

//UpdatePassword **
func (b *UsersServiceImpl) UpdatePassword(ctx context.Context,
	token string, updateRequest models.UpdatePasswordRequest) error {

	log.Logger(ctx).Info("UpdatePassword")

	if updateRequest.NewPassword != updateRequest.ConfirmPassword {
		return errors.New("New password and confirm password should be same")
	}

	err := b.userRepo.ValidateResetToken(ctx, token)

	if err != nil {
		return err
	}

	details, err := b.userRepo.GetEmailFromToken(ctx, token)

	if err != nil {
		return err
	}
	email := details.Email

	user, err := b.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(updateRequest.NewPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	err = b.userRepo.UpdatePassword(ctx, user.ID, string(passwordHash))

	if err != nil {
		return err
	}

	return nil
}

//GetUserCourseSections give a call to sections repository service
func (b *UsersServiceImpl) GetUserCourseSections(ctx context.Context,
	req models.GetUserCourseSectionsReq) (interface{}, error) {

	log.Logger(ctx).Info("GetUserCourseSections", req)

	id, err := b.userRepo.GetUserIdFromUUID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	req.ID = string(id)
	log.Logger(ctx).Info(id)
	resp, err := b.sectionCallRepo.GetSectionsInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getServiceBasepath() string {
	port := "8080"
	host := "localhost"
	basepathStr := fmt.Sprintf("%s:%s",
		host,
		port,
	)

	return basepathStr
}
