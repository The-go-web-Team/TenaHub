package repository

var a = 1

// import (
// 	"github.com/jinzhu/gorm"
// 	"github.com/TenaHub/api/entity"
// 	"github.com/TenaHub/api/user"
// )

// type UserGormRepo struct {
// 	conn *gorm.DB
// }

// func NewUserGormRepo(db *gorm.DB) user.UserRepository{
// 	return &UserGormRepo{conn:db}
// }

// func (adm *UserGormRepo) User(id uint) (*entity.User, []error) {
// 	user := entity.User{}
// 	errs := adm.conn.First(&user, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &user, errs

// }
// func (adm *UserGormRepo) Users() ([]entity.User, []error) {
// 	var users []entity.User
// 	errs := adm.conn.Find(&users).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return users, errs

// }
// func (adm *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
// 	user, errs := adm.User(id)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	errs = adm.conn.Delete(user, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return user, errs
// }




