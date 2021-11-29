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

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepository: repo}
}

func (service *userService) Save(user entity.User) entity.User {
	service.userRepository.SaveUser(user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.userRepository.FindAllUsers()
}

func (service *userService) Update(user entity.User) {
	service.userRepository.UpdateUser(user)
}
