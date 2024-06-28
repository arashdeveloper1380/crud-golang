package user

import "crud-api/app/models"

type UserServiceInterface interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	FindUserById(id uint) ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
