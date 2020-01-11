// package service

// import (
// 	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
// 	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice"
// )

// // HcserviceService is implementation of service.ServicesService
// type HcserviceService struct {
// 	servRepo hcservice.ServiceRepository
// }

// // NewHcserviceService creates HcserviceService object
// func NewHcserviceService(repo hcservice.ServiceRepository) *HcserviceService {
// 	return &HcserviceService{servRepo: repo}
// }

// // Services returns all healthcenter services
// func (hs *HcserviceService) Services(id uint) ([]entity.Service, []error) {
// 	services, errs := hs.servRepo.Services(id)

// 	if len(errs) > 0{
// 		return nil, errs
// 	}

// 	return services, nil
// }

// // Service returns single healthcenter service
// func (hs *HcserviceService) Service(id uint) (*entity.Service, []error) {
// 	service, errs := hs.servRepo.Service(id)

// 	if len(errs) > 0{
// 		return nil, errs
// 	}

// 	return service, nil
// }

// // UpdateService updates single healthcenter service
// func (hs *HcserviceService) UpdateService(service *entity.Service) (*entity.Service, []error) {
// 	serv, errs := hs.servRepo.UpdateService(service)

// 	if len(errs) > 0{
// 		return nil, errs
// 	}

// 	return serv, nil
// }

// // DeleteService deletes single healthcenter service
// func (hs *HcserviceService) DeleteService(id uint) (*entity.Service, []error) {
// 	service, errs := hs.servRepo.DeleteService(id)

// 	if len(errs) > 0{
// 		return nil, errs
// 	}


// 	return service, nil
// }

// // StoreService stores single healthcenter service
// func (hs *HcserviceService) StoreService(service *entity.Service) (*entity.Service, []error) {
// 	serv, errs := hs.servRepo.StoreService(service)

// 	if len(errs) > 0{
// 		return nil, errs
// 	}

// 	return serv, nil
// }
