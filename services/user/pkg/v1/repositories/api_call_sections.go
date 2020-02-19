package repositories

import (
	"context"
	"fmt"
	"io/ioutil"

	log "gitlab.com/scalent/goxpert/logs"

	"net/http"

	"gitlab.com/scalent/goxpert/models"
)

//SectionServiceRepo  all methods to get data form section service
type SectionServiceRepo interface {
	GetSectionsInfo(context.Context, models.GetUserCourseSectionsReq) (interface{}, error)
}

//SectionServiceRepoImpl enclose all the dependancies to cal section microservice
type SectionServiceRepoImpl struct {
}

//NewSectionServiceRepoImpl creates dependancies for section service
func NewSectionServiceRepoImpl() SectionServiceRepo {
	return &SectionServiceRepoImpl{}
}

//GetSectionsInfo ****
func (svc SectionServiceRepoImpl) GetSectionsInfo(ctx context.Context,
	req models.GetUserCourseSectionsReq) (interface{}, error) {

	url := fmt.Sprintf("http://course:8080/v1/course/%s/section/%s/info?user_id=%s",
		fmt.Sprint(req.CourseID),
		fmt.Sprint(req.SectionID),
		fmt.Sprint(req.ID),
	)

	response, err := http.Get(url)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err

	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err

	}
	fmt.Println(string(data))
	return string(data), nil
}
