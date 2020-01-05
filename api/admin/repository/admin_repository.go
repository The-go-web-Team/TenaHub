package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/admin"
	"github.com/TenaHub/api/entity"
	"fmt"
)

type AdminGormRepo struct {
	conn *gorm.DB
}

func NewAdminGormRepo(db *gorm.DB) admin.AdminRepository{
	return &AdminGormRepo{conn:db}
}

func (adm *AdminGormRepo) Admin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := entity.Admin{}
	errs := adm.conn.Where("email = ? AND password = ?", adminData.Email, adminData.Password).First(&admin).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &admin, errs
}
func (adm *AdminGormRepo) AdminById(id uint) (*entity.Admin, []error) {
	admin := entity.Admin{}
	errs := adm.conn.First(&admin, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &admin, errs
}

func (adm *AdminGormRepo) UpdateAdmin(adminData *entity.Admin) (*entity.Admin, []error) {
	admin := adminData
	fmt.Println(admin)
	errs := adm.conn.Save(admin).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return admin, errs
}


