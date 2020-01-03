package admin
import "github.com/TenaHub/api/entity"

type AdminRepository interface {
	Admin(id uint)(*entity.Admin, []error)
	UpdateAdmin(user *entity.Admin) (*entity.Admin, []error)

}

