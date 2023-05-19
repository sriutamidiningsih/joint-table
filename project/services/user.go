package services

import (
	userModel "dummy/project/model"
	userRepository "dummy/project/repository"
)

type Service interface {
	FindUserAll() ([]userModel.User, error)
	FindOrders() ([]userModel.Orders, error)
	FindByUserId(ID int) (userModel.User, error)
	Create(join userModel.RequestOrders) (userModel.User, error)
}

type serviceuser struct {
	repositoryUser userRepository.UserRepository
}

func NewServiceUser(repositoryUser userRepository.UserRepository) *serviceuser {
	return &serviceuser{repositoryUser}
}

func (service *serviceuser) FindUserAll() ([]userModel.User, error) {
	user, err := service.repositoryUser.FindUserAll()
	return user, err
}

func (service *serviceuser) FindOrders() ([]userModel.Orders, error) {
	order, err := service.repositoryUser.FindOrders()
	return order, err
}

func (service *serviceuser) FindByUserId(ID int) (userModel.User, error) {
	users, err := service.repositoryUser.FindByUserId(ID)
	return users, err
}

func (service *serviceuser) Create(join userModel.RequestOrders) (userModel.User, error) {
	users := userModel.User{
		Name:      join.Name,
		Email:     join.Email,
		OrderList: []userModel.Orders{},
	}
	newJoins, err := service.repositoryUser.Create(users)
	return newJoins, err
}
