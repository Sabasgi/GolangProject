package services

import (
	"repogin/internal/models"
	repositories "repogin/internal/repositories/sql"
)

type UserServiceStruct struct {
	User repositories.UserRepo
}

func NewUserService(Repo repositories.UserRepo) *UserServiceStruct {
	return &UserServiceStruct{
		User: Repo,
	}
}

func (u *UserServiceStruct) CreateService(user models.UserModel) error {
	Error := u.User.Create(user)
	if Error != nil {
		// fmt.Println("ERROR")
		return Error
	}
	return nil
}

func (u *UserServiceStruct) UpdateService(user models.UserModel) error {
	Error := u.User.Update(user)
	if Error != nil {
		// fmt.Println("ERROR")
		return Error
	}
	return nil
}
func (u *UserServiceStruct) GetAll() ([]models.UserModel, error) {
	users, Error := u.User.GetAll()
	if Error != nil {
		// fmt.Println("ERROR")
		return users, Error
	}
	return users, nil
}
func (u *UserServiceStruct) GetOneService(user models.UserModel) (models.UserModel, error) {
	User, Error := u.User.GetOne(user)
	if Error != nil {
		// fmt.Println("ERROR")
		return User, Error
	}
	return User, nil
}
