package service

import (
	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/basheer-shahrour/gin-server/repository"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Update(user entity.User)
}

type service struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &service{userRepository: repo}
}

func (service *service) Save(user entity.User) entity.User {
	service.userRepository.Save(user)
	return user
}

func (service *service) FindAll() []entity.User {
	return service.userRepository.FindAll()
}

func (service *service) Update(user entity.User) {
	service.userRepository.Update(user)
}
