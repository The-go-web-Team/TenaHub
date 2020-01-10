package user

import "github.com/TenaHub/api/entity"

type UserService interface {
	User(id uint) (*entity.User, []error)
	Users() ([]entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
}
