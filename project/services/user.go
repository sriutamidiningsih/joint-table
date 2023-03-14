package services

import (
	userModel "dummy/project/model"
	userRepository "dummy/project/repository"
)

type Service interface {
	FindAll() ([]userModel.User, error)
	FindproductByUserId(ID int) (userModel.User, error)
	// Create(Join userModel.RequestJoin) (userModel.Orders, error)
}

type serviceUser struct {
	repositoryUser userRepository.UserRepository
}

func NewServiceUser(repositoryUser userRepository.UserRepository) *serviceUser {
	return &serviceUser{repositoryUser}
}

func (service *serviceUser) FindAll() ([]userModel.User, error) {
	user, err := service.repositoryUser.FindAll()
	return user, err
}

func (service *serviceUser) FindproductByUserId(ID int) (userModel.User, error) {
	users, err := service.repositoryUser.FindProductByUserId(ID)
	return users, err
}
