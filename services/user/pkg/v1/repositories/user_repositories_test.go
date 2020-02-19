package repositories_test

import (
	"context"
	"testing"

	"gitlab.com/scalent/goxpert/database"
	"gitlab.com/scalent/goxpert/models"
	cmd "gitlab.com/scalent/goxpert/services/user/cmd/service"
	repository "gitlab.com/scalent/goxpert/services/user/pkg/v1/repositories"
)

//lDBHost=&"localhost"
// cmd.MysqlDBName=
var (
	mysqlUser     = "root"
	mysqlPassword = "password"
	mysqlDbhost   = "localhost"
	mysqlDbPort   = "3307"
	mysqlDbName   = "goxpert"
	emailSecret   = "RcZ3tf5EDO9b0AFK"
)

func init() {
	cmd.MysqlDBUser = &mysqlUser
	cmd.MysqlDBPass = &mysqlPassword
	cmd.MysqlDBHost = &mysqlDbhost
	cmd.MysqlDBAddr = &mysqlDbPort
	cmd.MysqlDBName = &mysqlDbName
	cmd.EmailSecret = &emailSecret
}

func TestUserRepositoryImpl_Login(t *testing.T) {

	if err := cmd.ValidateFlags(); err != nil {
		t.Error(err)
	}

	conPool := database.NewMysqlConnection(cmd.CreateConnectionString())
	userRepositoryImpl := repository.NewUserRepositoryImpl(conPool)

	type args struct {
		ctx  context.Context
		user models.UserLoginRequest
	}
	tests := []struct {
		name               string
		userRepositoryImpl repository.UserRepository
		args               args
		wantErr            bool
	}{
		{
			"Test 1 :First Test ",
			userRepositoryImpl,
			args{
				context.Background(),
				models.UserLoginRequest{
					Email:    "rahul@gmail.com",
					Password: "password",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.userRepositoryImpl.Login(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
