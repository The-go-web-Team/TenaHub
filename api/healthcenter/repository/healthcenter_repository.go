package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/healthcenter"
	"fmt"
)

type HealthCenterGormRepo struct {
	conn *gorm.DB
}

func NewHealthCenterGormRepo(db *gorm.DB) healthcenter.HealthCenterRepository{
	return &HealthCenterGormRepo{conn:db}
}

func (adm HealthCenterGormRepo) HealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	errs := adm.conn.First(&healthcenter, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &healthcenter, errs
}
func (adm *HealthCenterGormRepo) HealthCenters() ([]entity.HealthCenter, []error) {
	var healthcenters []entity.HealthCenter
	errs := adm.conn.Find(&healthcenters).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenters, errs

}
func (adm *HealthCenterGormRepo) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := adm.HealthCenter(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = adm.conn.Delete(healthcenter, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}

func (adm *HealthCenterGormRepo) UpdateHealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := healthcenterData
	fmt.Println(healthcenter)
	errs := adm.conn.Save(healthcenter).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, errs
}





