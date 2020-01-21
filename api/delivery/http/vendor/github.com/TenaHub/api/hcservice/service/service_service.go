package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/hcservice"
)

type ServiceService struct {
	serviceRepo hcservice.ServiceRepository
}
func NewServiceService(serv hcservice.ServicesService)(service ServiceService){
	return ServiceService{serviceRepo:serv}
}

func (adm *ServiceService) Service(id uint) (*entity.Service, []error){
	service, errs := adm.serviceRepo.Service(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}
func (adm *ServiceService) PendingService() ([]entity.Service, []error) {
	services, errs := adm.serviceRepo.PendingService()
	if len(errs) > 0 {
		return nil, errs
	}
	return services, errs
}
func (adm *ServiceService) Services(id uint) ([]entity.Service, []error) {
	healthCenters, errs := adm.serviceRepo.Services(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return healthCenters, errs
}
func (adm *ServiceService) UpdateService(serviceData *entity.Service) (*entity.Service, []error) {
	service, errs := adm.serviceRepo.UpdateService(serviceData)
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}
func (adm *ServiceService) StoreService(service *entity.Service) (*entity.Service, []error) {
	service, errs := adm.serviceRepo.StoreService(service)
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}
func (adm *ServiceService) DeleteService(id uint) (*entity.Service, []error) {
	service, errs := adm.serviceRepo.DeleteService(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return service, errs
}


