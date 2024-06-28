package user

import (
	"crud-api/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func Instance(db *gorm.DB) UserInterface {
	return &UserRepository{db}
}

func (userRepository *UserRepository) CreateUser(user *models.User) error {
	return userRepository.db.Create(user).Error
}

func (userRepository *UserRepository) GetAllUsers() ([]models.User, error) {
	var user []models.User
	err := userRepository.db.Find(&user).Error

	return user, err

}

func (userRepository *UserRepository) FindUserById(id uint) (models.User, error) {
	var user models.User
	err := userRepository.db.First(&user, id).Error

	return user, err
}

func (userRepository *UserRepository) UpdateUser(user *models.User) error {
	return userRepository.db.Save(user).Error
}

func (userRepository *UserRepository) DeleteUser(id uint) error {
	return userRepository.db.Delete(&models.User{}, id).Error
}
