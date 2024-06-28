package user

import (
	"crud-api/app/models"
	"crud-api/app/repositories/user"
)

type UserService struct {
	repo user.UserRepository
}

func Instance(repo user.UserInterface) UserService {
	return &user.UserRepository{repo}
}

func (UserService *UserService) CreateUser(user *models.User) error {
	return UserService.repo.CreateUser(user)
}

func (UserSevice *UserService) GetAllUsers() ([]models.User, error) {
	return UserSevice.repo.GetAllUsers()
}

func (UserSevice *UserService) FindUserById(id uint) ([]models.User, error) {
	return UserSevice.repo.FindUserById(id)
}

func (UserSevice *UserService) UpdateUser(user *models.User) error {
	return UserSevice.repo.UpdateUser(user)
}

func (UserSevice *UserService) DeleteUser(id uint) error {
	return UserSevice.repo.DeleteUser(id)
}
