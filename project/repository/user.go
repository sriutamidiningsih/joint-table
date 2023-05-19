package repository

import (
	userModel "dummy/project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserAll() ([]userModel.User, error)
	FindOrders() ([]userModel.Orders, error)
	FindByUserId(ID int) (userModel.User, error)
	Create(user userModel.User) (userModel.User, error)
}

type repositoryuser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryuser {
	return &repositoryuser{db}
}

func (repositoryuser *repositoryuser) FindUserAll() ([]userModel.User, error) {
	var user []userModel.User
	err := repositoryuser.db.Joins("Order").Select("*").Find(&user).Error
	return user, err
}

func (repositoryuser *repositoryuser) FindOrders() ([]userModel.Orders, error) {
	var order []userModel.Orders
	err := repositoryuser.db.Find(&order).Error
	return order, err
}

func (repositoryuser *repositoryuser) FindByUserId(ID int) (userModel.User, error) {
	var join userModel.User
	err := repositoryuser.db.Preload("Order").Where("id=?", ID).Find(&join).Error
	return join, err
}

func (repositoryuser *repositoryuser) Create(user userModel.User) (userModel.User, error) {
	err := repositoryuser.db.Create(&user).Error
	return user, err
}
