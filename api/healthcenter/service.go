package healthcenter

import "github.com/TenaHub/api/entity"

type HealthCenterService interface {
	HealthCenterById(id uint) (*entity.HealthCenter, []error)
	HealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)
	HealthCenters() ([]entity.HealthCenter, []error)
	DeleteHealthCenter(id uint) (*entity.HealthCenter, []error)
	UpdateHealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)

}