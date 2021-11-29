package repository

import (
	"github.com/basheer-shahrour/gin-server/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(user entity.User)
	UpdateUser(user entity.User)
	DeleteUser(user entity.User)
	FindAllUsers() []entity.User
}

type userRepository struct {
	dbConnection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{dbConnection: db}
}

func (db *userRepository) SaveUser(user entity.User) {
	db.dbConnection.Create(&user)
}

func (db *userRepository) UpdateUser(user entity.User) {
	db.dbConnection.Save(&user)
}

func (db *userRepository) DeleteUser(user entity.User) {
	db.dbConnection.Delete(&user)
}

func (db *userRepository) FindAllUsers() []entity.User {
	var users []entity.User
	db.dbConnection.Set("gorm:auto_preload", true).Find(&users)
	return users
}
