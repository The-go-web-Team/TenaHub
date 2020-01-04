package service

import "github.com/TenaHub/api/entity"

type ServiceRepository interface {
	Service(id uint) (*entity.Service, []error)
	PendingService() ([]entity.Service, []error)
	Services() ([]entity.Service, []error)
	UpdateService(user *entity.Service) (*entity.Service, []error)
	StoreService(user *entity.Service) (*entity.Service, []error)
	DeleteService(id uint) (*entity.Service, []error)
}