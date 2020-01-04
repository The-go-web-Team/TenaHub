package healthcenter

import "github.com/TenaHub/api/entity"

type HealthCenterRepository interface {
	HealthCenter(id uint) (*entity.HealthCenter, []error)
	HealthCenters() ([]entity.HealthCenter, []error)
	DeleteHealthCenter(id uint) (*entity.HealthCenter, []error)
	UpdateHealthCenter(healthcenter *entity.HealthCenter) (*entity.HealthCenter, []error)

}