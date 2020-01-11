package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/healthcenter"
)

type HealthCenterService struct {
	healthCenterRepo healthcenter.HealthCenterRepository
}
func NewHealthCenterService(serv healthcenter.HealthCenterService)(admin *HealthCenterService){
	return &HealthCenterService{healthCenterRepo:serv}
}


func (adm *HealthCenterService) HealthCenterById(id uint) (*entity.HealthCenter, []error) {
	healthCenter, errs := adm.healthCenterRepo.HealthCenterById(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenter, errs
}
func (adm *HealthCenterService) HealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthCenter, errs := adm.healthCenterRepo.HealthCenter(healthcenter)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenter, errs
}
func (adm *HealthCenterService) HealthCenters() ([]entity.HealthCenter, []error) {
	healthCenters, errs := adm.healthCenterRepo.HealthCenters()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenters, errs
}
func (adm *HealthCenterService) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.healthCenterRepo.DeleteHealthCenter(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}
func (adm *HealthCenterService) UpdateHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.healthCenterRepo.UpdateHealthCenter(healthcenterData)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}
