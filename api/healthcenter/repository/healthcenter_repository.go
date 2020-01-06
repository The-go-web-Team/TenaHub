package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/healthcenter"
	"github.com/TenaHub/api/delivery/http/handler"
)

type HealthCenterGormRepo struct {
	conn *gorm.DB
}

func NewHealthCenterGormRepo(db *gorm.DB) healthcenter.HealthCenterRepository{
	return &HealthCenterGormRepo{conn:db}
}

func (adm HealthCenterGormRepo) HealthCenterById(id uint) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	errs := adm.conn.First(&healthcenter, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &healthcenter, errs
}
func (adm HealthCenterGormRepo) HealthCenter(healthcenterData *entity.HealthCenter) (*entity.HealthCenter, []error) {
	healthcenter := entity.HealthCenter{}
	//errs := adm.conn.Where("email = ? AND password = ?", healthcenterData.Email, healthcenterData.Password).First(&healthcenter).GetErrors()
	//if len(errs) > 0 {
	//	return nil, errs
	//}
	errs := adm.conn.Select("password").Where("email = ? ", healthcenterData.Email).First(&healthcenter).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	same := handler.VerifyPassword(healthcenterData.Password, healthcenter.Password)
	if same {
		errs := adm.conn.Where("email = ?", healthcenterData.Email).First(&healthcenter).GetErrors()
		return &healthcenter, errs
	}
	return nil, errs
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
	healthcenter, errs := adm.HealthCenterById(id)
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
	data := entity.HealthCenter{}
	healthcenter.Password,_ = handler.HashPassword(healthcenterData.Password)
	//errs := adm.conn.Save(healthcenter).GetErrors()
	errs := adm.conn.Model(&data).Updates(healthcenter).Error
	if errs != nil {
		return nil, []error{errs}
	}
	return healthcenter, []error{errs}
}





