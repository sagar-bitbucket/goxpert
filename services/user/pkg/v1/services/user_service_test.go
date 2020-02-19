package services_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.com/scalent/goxpert/email"
	"gitlab.com/scalent/goxpert/models"
	mock "gitlab.com/scalent/goxpert/services/user/mocks"
	service "gitlab.com/scalent/goxpert/services/user/pkg/v1/services"
)

func TestUsersServiceImpl_CreateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var emailSecret = os.Getenv("EMAIL_SECRET")
	email := email.NewSendInBlue(emailSecret)

	mockUsersRepo := mock.NewMockUserRepository(ctrl)
	usersService := service.NewUserServiceImpl(mockUsersRepo, email)

	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name             string
		args             args
		emailExpect      func()
		createUserExpect func()
		wantErr          bool
	}{
		{
			"Test 1 : Create Users",
			args{
				context.Background(),
				models.User{},
			},
			func() {
				mockUsersRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(models.User{}, nil)
			},
			func() {

				mockUsersRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return("", nil)
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.emailExpect()

			_, err := usersService.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersServiceImpl.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//fmt.Println("GOT", got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UsersServiceImpl.CreateUser() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUsersServiceImpl_Login(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var emailSecret = os.Getenv("EMAIL_SECRET")
	email := email.NewSendInBlue(emailSecret)

	mockUsersRepo := mock.NewMockUserRepository(ctrl)
	usersService := service.NewUserServiceImpl(mockUsersRepo, email)

	type args struct {
		ctx  context.Context
		user models.UserLoginRequest
	}
	tests := []struct {
		name    string
		args    args
		fun     func()
		wantErr bool
	}{
		{
			"Test 1: Login Test",
			args{
				context.Background(),
				models.UserLoginRequest{},
			},
			func() {
				mockUsersRepo.EXPECT().Login(context.Background(), gomock.Any()).Return(models.User{}, nil)
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fun()
			gotResp, err := usersService.Login(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsersServiceImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotResp)
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("UsersServiceImpl.Login() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}
