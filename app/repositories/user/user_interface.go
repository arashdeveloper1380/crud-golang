package user

import "crud-api/app/models"

type UserInterface interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	FindUserById(id uint) ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
