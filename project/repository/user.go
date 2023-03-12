package repository

import (
	userModel "dummy/project/model"

	"gorm.io/gorm"
)

type UserRepository interface{
	FindAll() ([]userModel.User, error)
	Create(User userModel.User)(userModel.User, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

func (repositoryUser *repositoryUser) FindAll() ([]userModel.User, error) {
	var user []userModel.User
	err := repositoryUser.db.Joins("UserAdress").Select("*").Find(&user).Error
	return user, err
}

