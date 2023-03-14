package repository

import (
	userModel "dummy/project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]userModel.User, error)
	FindProductByUserId(ID int) (userModel.User, error)
	Create(Join userModel.RequestJoin) (userModel.RequestJoin, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

func (repositoryUser *repositoryUser) FindAll() ([]userModel.User, error) {
	var user []userModel.User
	err := repositoryUser.db.Joins("Order").Select("*").Find(&user).Error
	return user, err
}

func (repositoryUser *repositoryUser) FindProductByUserId(ID int) (userModel.User, error) {
	var join userModel.User
	err := repositoryUser.db.Preload("Order").Where("id=?", ID).Find(&join).Error
	return join, err
}
