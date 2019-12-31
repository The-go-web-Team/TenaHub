package repository

import (
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
)

// ServiceGormRepository is repository implements service.ServiceRepository
type ServiceGormRepository struct {
	conn *gorm.DB
}

// NewServiceGormRepo creates and returns new ServiceGormRepo object
func NewServiceGormRepo(conn *gorm.DB) *ServiceGormRepository {
	return &ServiceGormRepository{conn: conn}
}

// Services returns all services from the database
func (sr *ServiceGormRepository) Services(id uint)([]entity.Service, []error) {
	services := []entity.Service{}
	errs := sr.conn.Where("health_center_id = ?", id).Find(&services).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return services , nil
}

// Service returns single service from database
func (sr *ServiceGormRepository) Service(id uint)(*entity.Service, []error) {
	service := entity.Service{}
	errs := sr.conn.First(&service, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &service , nil
}

// UpdateService updates single service on the database
func (sr *ServiceGormRepository) UpdateService(service *entity.Service)(*entity.Service, []error) {
	serv := service
	errs := sr.conn.Save(&serv).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return serv , nil
}

// DeleteService deletes single service from the database
func (sr *ServiceGormRepository) DeleteService(id uint)(*entity.Service, []error) {
	serv, errs := sr.Service(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = sr.conn.Delete(serv, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return serv , nil
}

// StoreService stores single service to the database
func (sr *ServiceGormRepository) StoreService(service *entity.Service)(*entity.Service, []error) {
	serv := service
	errs := sr.conn.Create(serv).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return serv , nil
}