package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "gitlab.com/scalent/goxpert/logs"
	"gitlab.com/scalent/goxpert/models"
)

//UserServiceRepo  all methods to get data form User service
type UserServiceRepo interface {
	GetUsersInfoByUUID(context.Context, string) (*models.User, error)
}

//UserServiceRepoImpl enclose all the dependancies to cal User microservice
type UserServiceRepoImpl struct {
}

//NewUserServiceRepoImpl creates dependancies for User service
func NewUserServiceRepoImpl() UserServiceRepo {
	return &UserServiceRepoImpl{}
}

//GetUsersInfoByUUID ****
func (svc UserServiceRepoImpl) GetUsersInfoByUUID(ctx context.Context, uuid string) (*models.User, error) {

	url := fmt.Sprintf("http://user:8080/user/%s",
		fmt.Sprint(uuid),
	)

	response, err := http.Get(url)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err

	}
	decoder := json.NewDecoder(response.Body)
	user := models.User{}
	err = decoder.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
