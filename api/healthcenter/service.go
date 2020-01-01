package healthcenter

import "github.com/TenaHub/api/entity"

type HealthCenterService interface {

	HealthCenter(id uint) (*entity.HealthCenter, []error)
	HealthCenters() ([]entity.HealthCenter, []error)
	DeleteHealthCenter(id uint) (*entity.HealthCenter, []error)
}