package user

import "github.com/TenaHub/api/entity"

type UserRepository interface {
	User(id uint) (*entity.User, []error)
	Users() ([]entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
}