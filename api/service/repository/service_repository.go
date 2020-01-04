package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/service"
)

type ServiceGormRepo struct {
	conn *gorm.DB
}

func NewServiceGormRepo(db *gorm.DB) service.ServiceRepository{
	return &ServiceGormRepo{conn:db}
}
func (adm *ServiceGormRepo) Service(id uint) (*entity.Service, []error) {
	var service entity.Service
	errs := adm.conn.First(&service, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &service, errs
}
func (adm *ServiceGormRepo) PendingService() ([]entity.Service, []error) {
	var services []entity.Service
	errs := adm.conn.Find(&services).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return services, errs
}
func (adm *ServiceGormRepo) Services() ([]entity.Service, []error) {
	var services []entity.Service
	errs := adm.conn.Find(&services).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return services, errs
}
func (adm *ServiceGormRepo) UpdateService(serviceData *entity.Service) (*entity.Service, []error) {
	service := serviceData
	errs := adm.conn.Save(service).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs

}
func (adm *ServiceGormRepo) StoreService(serviceData *entity.Service) (*entity.Service, []error) {
	service := serviceData
	errs := adm.conn.Create(service).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}
func (adm *ServiceGormRepo) DeleteService(id uint) (*entity.Service, []error) {
	service, errs := adm.Service(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = adm.conn.Delete(service, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}
