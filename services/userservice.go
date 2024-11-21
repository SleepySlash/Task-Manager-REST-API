package services

import "Task-Manager-REST-API/model"

type UserService interface {
	newUser(name string) (model.User, error)
	getUser(name string) (model.User, error)
	getAllUser() ([]model.User, error)
	deleteUser(name string) (model.User, error)
	deleteAllUser() error
	updateUser(name string) (model.User, error)
}
type userService struct {
	repo model.Users
}

func NewUserService(repo model.Users) UserService {
	return &userService{
		repo: repo,
	}
}
func (t *userService) newUser(name string) (model.User, error)    {}
func (t *userService) getUser(name string) (model.User, error)    {}
func (t *userService) getAllUser() ([]model.User, error)          {}
func (t *userService) deleteUser(name string) (model.User, error) {}
func (t *userService) deleteAllUser() error                       {}
func (t *userService) updateUser(name string) (model.User, error) {}
