package services

import (
	userModel "dummy/project/model"
	userRepository "dummy/project/repository"
)

type Service interface {
	FindUserAll() ([]userModel.Users, error)
	FindJoins() ([]userModel.Users, error)
	FindByUserId(ID int) (userModel.Users, error)
	Create(join userModel.RequestOrders) (userModel.Users, error)
}

type serviceuser struct {
	repositoryUser userRepository.UserRepository
}

func NewServiceUser(repositoryUser userRepository.UserRepository) *serviceuser {
	return &serviceuser{repositoryUser}
}

func (service *serviceuser) FindUserAll() ([]userModel.Users, error) {
	user, err := service.repositoryUser.FindUserAll()
	return user, err
}

func (service *serviceuser) FindJoins() ([]userModel.Users, error) {
	order, err := service.repositoryUser.FindJoin()
	return order, err
}

func (service *serviceuser) FindByUserId(ID int) (userModel.Users, error) {
	users, err := service.repositoryUser.FindByUserId(ID)
	return users, err
}

func (service *serviceuser) Create(join userModel.RequestOrders) (userModel.Users, error) {
	users := userModel.Users{
		Name:  join.Name,
		Email: join.Email,
		OrderList: []userModel.Orders{
			{
				NameProduct: join.NameProduct,
				TotalOrder:  join.TotalOrder,
			},
		},
	}
	newJoins, err := service.repositoryUser.Create(users)
	return newJoins, err
}
