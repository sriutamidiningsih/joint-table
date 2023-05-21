package repository

import (
	userModel "dummy/project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserAll() ([]userModel.Users, error)
	FindOrders() ([]userModel.Orders, error)
	FindByUserId(ID int) (userModel.Users, error)
	Create(user userModel.Users) (userModel.Users, error)
}

type repositoryuser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryuser {
	return &repositoryuser{db}
}

func (repositoryuser *repositoryuser) FindUserAll() ([]userModel.Users, error) {
	var user []userModel.Users
	err := repositoryuser.db.Find(&user).Error
	return user, err
}

func (repositoryuser *repositoryuser) FindOrders() ([]userModel.Orders, error) {
	var order []userModel.Orders
	err := repositoryuser.db.Table("order").Joins("JOIN users ON order.id_user = users.id").
		Select("order.total_order, order.name_product, users.name, users.email").Find(&order).Error
	return order, err
}

func (repositoryuser *repositoryuser) FindByUserId(ID int) (userModel.Users, error) {
	var join userModel.Users
	err := repositoryuser.db.Preload("OrderList").Where("users.id=?", ID).Find(&join).Error
	return join, err
}

func (repositoryuser *repositoryuser) Create(user userModel.Users) (userModel.Users, error) {
	err := repositoryuser.db.Create(&user).Error
	return user, err
}
