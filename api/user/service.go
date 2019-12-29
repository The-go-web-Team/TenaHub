package user

import "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

// UserService is
type UserService interface {
	Users() ([]entity.User, []error)
	User(user *entity.User) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id int) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}
