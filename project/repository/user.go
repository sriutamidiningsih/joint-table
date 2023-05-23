package repository

import (
	userModel "dummy/project/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserAll() ([]userModel.Users, error)
	FindJoin() ([]userModel.Users, error)
	FindByUserId(ID int) (userModel.Users, error)
	Create(user userModel.Users) (userModel.Users, error)
	CreateOrders(orders userModel.Orders) (userModel.Orders, error)
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

func (repositoryuser *repositoryuser) FindJoin() ([]userModel.Users, error) {
	var joins []userModel.Users
	err := repositoryuser.db.Preload("OrderList").Find(&joins).Error
	//err := repositoryuser.db.Table("order").Joins("JOIN users ON order.id_user = users.id").Find(&order).Error
	return joins, err
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

func (repositoryuser *repositoryuser) CreateOrders(orders userModel.Orders) (userModel.Orders, error) {
	err := repositoryuser.db.Create(&orders).Error
	return orders, err
}
