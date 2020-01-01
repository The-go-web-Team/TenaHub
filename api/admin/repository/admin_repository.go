package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/admin"
	"github.com/TenaHub/api/entity"
)

type AdminGormRepo struct {
	conn *gorm.DB
}

func NewAdminGormRepo(db *gorm.DB) admin.AdminRepository{
	return &AdminGormRepo{conn:db}
}

func (adm *AdminGormRepo) Admin(id uint) (*entity.Admin, []error) {
	admin := entity.Admin{}
	errs := adm.conn.First(&admin, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &admin, errs
}



