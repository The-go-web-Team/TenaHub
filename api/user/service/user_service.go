package service

import (
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/user"
)

type UserService struct {
	userRepo user.UserRepository
}
func NewUserService(serv user.UserRepository)(admin *UserService){
	return &UserService{userRepo:serv}
}


func (adm *UserService) User(id uint) (*entity.User, []error) {
	user, errs := adm.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return user, errs
}
func (adm *UserService) Users() ([]entity.User, []error) {
	users, errs := adm.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return users, errs
}
func (adm *UserService) DeleteUser(id uint) (*entity.User, []error) {
	user, errs := adm.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return user, errs
}
