package healthcenter

import "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

// HealthCenterService is
type HealthCenterService interface {
	HealthCenter(id int) (*entity.HealthCenter, []error)
	HealthCenters() ([]entity.HealthCenter, []error)
	UpdateHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error)
	StoreHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error)
	DeleteHealthCenter(id int) (*entity.HealthCenter, []error)
}
