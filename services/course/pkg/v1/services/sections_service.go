package services

import (
	"context"

	log "gitlab.com/scalent/goxpert/logs"

	"gitlab.com/scalent/goxpert/models"

	repository "gitlab.com/scalent/goxpert/services/course/pkg/v1/repositories"
)

// SectionsService describes the service.
type SectionsService interface {
	CreateSection(context.Context, models.Sections) (interface{}, error)
	GetSections(context.Context, uint32, uint32) (*models.Sections, error)
	GetAllSections(context.Context, uint32) (*[]models.Sections, error)
	UpdateSections(context.Context, models.Sections) (interface{}, error)
	UpdateSectionsSequence(context.Context, models.ChnageSectionSequence) (interface{}, error)
	GetSectionsInfo(context.Context, models.SectionInfoReq) (interface{}, error)
}

//SectionsServiceImpl **
type SectionsServiceImpl struct {
	SectionsRepo repository.SectionsRepository
}

//NewSectionsServiceImpl inject depedancies user repositiory
func NewSectionsServiceImpl(SectionsRepo repository.SectionsRepository) SectionsService {
	return &SectionsServiceImpl{SectionsRepo: SectionsRepo}
}

// CreateSection logs CreateSection Request and send it to repository layer
func (secSvcIml SectionsServiceImpl) CreateSection(ctx context.Context,
	sectionReq models.Sections) (interface{}, error) {

	log.Logger(ctx).Info("Create Section Request ", sectionReq)

	resp, err := secSvcIml.SectionsRepo.CreateSection(ctx, sectionReq)

	return resp, err
}

//GetSections *
func (secSvcIml SectionsServiceImpl) GetSections(ctx context.Context,
	courseID uint32, sectionID uint32) (*models.Sections, error) {

	log.Logger(ctx).Info("Get Section By ID :  ", "courseID : ", courseID, "sectionID : ", sectionID)

	resp, err := secSvcIml.SectionsRepo.GetSections(ctx, courseID, sectionID)

	return resp, err

}

//GetAllSections returns sections from database
func (secSvcIml SectionsServiceImpl) GetAllSections(ctx context.Context, page uint32) (*[]models.Sections, error) {
	log.Logger(ctx).Info("Get All Sections :")

	resp, err := secSvcIml.SectionsRepo.GetAllSections(ctx, page)

	return resp, err
}

//UpdateSections updated
func (secSvcIml SectionsServiceImpl) UpdateSections(ctx context.Context,
	section models.Sections) (interface{}, error) {
	log.Logger(ctx).Info("UpdateSections:", section)
	resp, err := secSvcIml.SectionsRepo.UpdateSections(ctx, section)
	return resp, err

}

//UpdateSectionsSequence log request and forward to Repository
func (secSvcIml SectionsServiceImpl) UpdateSectionsSequence(ctx context.Context,
	req models.ChnageSectionSequence) (interface{}, error) {

	log.Logger(ctx).Info("UpdateSectionsSequence:", req)
	resp, err := secSvcIml.SectionsRepo.UpdateSectionsSequence(ctx, req)
	return resp, err

}

//GetSectionsInfo log request and forward to Repository
func (secSvcIml SectionsServiceImpl) GetSectionsInfo(ctx context.Context,
	req models.SectionInfoReq) (interface{}, error) {

	log.Logger(ctx).Info("GetSectionsInfo:", req)
	resp, err := secSvcIml.SectionsRepo.GetSectionsInfo(ctx, req)
	return resp, err

}
