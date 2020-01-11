package healthcenter

import "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

// HealthCenterService is
type HealthCenterService interface {
	HealthCenter(id uint) (*entity.HealthCenter, []error)
	HealthCenters(value string, column string) ([]entity.Hcrating, []error)
	Top(amount uint)([]entity.Hcrating, []error)
	UpdateHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error)
	StoreHealthCenter(hc entity.HealthCenter) (*entity.HealthCenter, []error)
	DeleteHealthCenter(id uint) (*entity.HealthCenter, []error)
}
