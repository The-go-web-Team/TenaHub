package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/hcservice"
	"errors"
)

type MockServiceGormRepo struct {
	conn *gorm.DB
}

func NewMockServiceGormRepo(db *gorm.DB) hcservice.ServiceRepository{
	return &MockServiceGormRepo{conn:db}
}
func (adm *MockServiceGormRepo) Service(id uint) (*entity.Service, []error) {
	var service entity.Service
	service = entity.MockService
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &service, nil
}
func (adm *MockServiceGormRepo) PendingService() ([]entity.Service, []error) {
	var services []entity.Service
	services = append(services,entity.MockService, entity.MockService)
	return services, nil
}
func (adm *MockServiceGormRepo) Services(id uint) ([]entity.Service, []error) {
	var services []entity.Service
	services = append(services,entity.MockService, entity.MockService)
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return services, nil
}
func (adm *MockServiceGormRepo) UpdateService(serviceData *entity.Service) (*entity.Service, []error) {
	service := serviceData

	return service, nil

}
func (adm *MockServiceGormRepo) StoreService(serviceData *entity.Service) (*entity.Service, []error) {
	service := serviceData

	return service, nil
}
func (adm *MockServiceGormRepo) DeleteService(id uint) (*entity.Service, []error) {
	service, errs := adm.Service(id)
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return service, errs
}
