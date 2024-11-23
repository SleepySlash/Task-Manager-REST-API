package services

import (
	"Task-Manager-REST-API/middleware"
	"Task-Manager-REST-API/model"
	"log"
	"time"
)

type UserService interface {
	NewUser(user model.User) error
	GetUser(user model.User) (string, error)
	GetAllUser() ([]model.User, error)
	DeleteUser(name string) (string, error)
	DeleteAllUser() error
	UpdateUser(name string, user model.User) (bool, error)
}
type userService struct {
	repo model.Users
}

func NewUserService(repo model.Users) UserService {
	return &userService{
		repo: repo,
	}
}
func (t *userService) NewUser(user model.User) error {
	currentDate := time.Now().Format("20060102")
	userid := user.Username + currentDate
	user.UserId = userid
	err := t.repo.Create(user)
	if err != nil {
		log.Fatal("error during new user creation")
		return err
	}
	return nil
}
func (t *userService) GetUser(user model.User) (string, error) {
	res, err := t.repo.Get(user)
	if err != nil {
		log.Fatal("error while fetching user")
		return "", err
	}

	tokenString, err := middleware.CreateToken(res.UserId)
	if err != nil {
		log.Fatal("Error during the creation of the jwt token")
		return "", err
	}
	return tokenString, nil
}
func (t *userService) GetAllUser() ([]model.User, error) {
	var result []model.User
	result, err := t.repo.All()
	if err != nil {
		log.Fatal("error while fetching all the existing user")
		return result, err
	}
	return result, nil
}
func (t *userService) DeleteUser(userid string) (string, error) {
	err := t.repo.Delete(userid)
	if err != nil {
		log.Fatal("error while fetching all the existing user")
		return "", err
	}
	return userid, nil

}
func (t *userService) DeleteAllUser() error {
	err := t.repo.DeleteAll()
	if err != nil {
		log.Fatal("error while deleting all the existing user")
		return err
	}
	return nil
}
func (t *userService) UpdateUser(userid string, user model.User) (bool, error) {
	_, err := t.repo.Update(userid, user)
	if err != nil {
		log.Fatal("error while updating the user")
		return false, err
	}
	return true, nil
}
