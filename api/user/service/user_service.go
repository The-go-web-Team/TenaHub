package service

import (
	"fmt"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user"
)

// UserService implements interface user.UserService
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService creates and returns UserService object
func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// Users returns all users from the database
func (us *UserService) Users() ([]entity.User, []error) {
	users, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}

	return users, nil
}

// User returns a single users from the database with user name and password
func (us *UserService) User(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.User(user)
	fmt.Println(errs)
	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// UserByID returns a single user by its id from the database
func (us *UserService) UserByID(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.UserByID(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// UpdateUser updates user from the database
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// DeleteUser deletes a single user from the database by it's id
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// StoreUser will insert a new user to the database
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}
