package services

import "Task-Manager-REST-API/model"

type UserService interface {
	NewUser(user model.User) (model.User, error)
	GetUser(user model.User) (bool, error)
	GetAllUser() ([]model.User, error)
	DeleteUser(name string) (model.User, error)
	DeleteAllUser() error
	UpdateUser(name string, user model.User) (model.User, error)
}
type userService struct {
	repo model.Users
}

func NewUserService(repo model.Users) UserService {
	return &userService{
		repo: repo,
	}
}
func (t *userService) NewUser(user model.User) (model.User, error)                 {}
func (t *userService) GetUser(user model.User) (bool, error)                       {}
func (t *userService) GetAllUser() ([]model.User, error)                           {}
func (t *userService) DeleteUser(name string) (model.User, error)                  {}
func (t *userService) DeleteAllUser() error                                        {}
func (t *userService) UpdateUser(name string, user model.User) (model.User, error) {}
