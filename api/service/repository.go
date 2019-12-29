package service

import (
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
)

// ServiceRepository is
type ServiceRepository interface {
	Services()([]entity.Service, []error)
	Service(id int)(*entity.Service, []error)
	UpdateService(service *entity.Service)(*entity.Service, []error)
	DeleteService(id int)(*entity.Service, []error)
	StoreService(service *entity.Service)(*entity.Service, []error)
}