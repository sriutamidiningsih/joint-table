package services

import (
	userModel "dummy/project/model"
	userRepository "dummy/project/repository"
)

type Service interface {
	FindAll() ([]userModel.User, error)
	Create(User userModel.User) (userModel.User, error)
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
