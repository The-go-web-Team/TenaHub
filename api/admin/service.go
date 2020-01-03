package admin

import "github.com/TenaHub/api/entity"

type AdminService interface {
	Admin(id uint)(*entity.Admin, []error)
	UpdateAdmin(user *entity.Admin) (*entity.Admin, []error)

}
