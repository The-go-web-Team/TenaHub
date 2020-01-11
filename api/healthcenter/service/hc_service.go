package service

import (
	"fmt"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/healthcenter"
)

// HcService implements healthcenter.HealthCenterService
type HcService struct {
	hcRepo healthcenter.HealthCenterRepository
}

// NewHcService creates object of HcService
func NewHcService(repo healthcenter.HealthCenterRepository) *HcService {
	return &HcService{hcRepo: repo}
}

// HealthCenter returns single healthcenter data
func (hcs *HcService) HealthCenter(id uint) (*entity.HealthCenter, []error) {
	healthcenter, errs := hcs.hcRepo.HealthCenter(uint(id))

	if len(errs) > 0 {
		return nil, errs
	}
	return healthcenter, nil
}

// HealthCenters returns all healthcenters data
func (hcs *HcService) HealthCenters(value string, column string) ([]entity.Hcrating, []error) {
	healthcenters, errs := hcs.hcRepo.HealthCenters(value, column)

	if errs != nil {
		return nil, errs
	}
	return healthcenters, nil
}

// Top returns healthcenters with rating from database
func (hcs *HcService) Top(amount uint) ([]entity.Hcrating, []error) {
	result, errs := hcs.hcRepo.Top(amount)
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println(result)
	return result, nil
}

// SearchHealthCenter searches healthcenters by name
// func (hcs *HcService) SearchHealthCenter(name string)([]entity.HealthCenter, []error) {
// return nil, nil
// }

// UpdateHealthCenter updates healthcenter
func (hcs *HcService) UpdateHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error) {
	return nil, nil
}

// StoreHealthCenter stores healthcenter data
func (hcs *HcService) StoreHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error) {
	return nil, nil
}

// DeleteHealthCenter deletes healthcenter
func (hcs *HcService) DeleteHealthCenter(id uint) (*entity.HealthCenter, []error) {
	return nil, nil
}
