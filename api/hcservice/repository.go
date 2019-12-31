package hcservice

import (
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
)

// ServiceRepository is
type ServiceRepository interface {
	Services(id uint)([]entity.Service, []error)
	Service(id uint)(*entity.Service, []error)
	UpdateService(service *entity.Service)(*entity.Service, []error)
	DeleteService(id uint)(*entity.Service, []error)
	StoreService(service *entity.Service)(*entity.Service, []error)
}