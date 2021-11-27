package repository

import (
	"github.com/basheer-shahrour/gin-server/database"
	"github.com/basheer-shahrour/gin-server/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(user entity.User)
	FindAll() []entity.User
	CloseDB()
}

type databaseConnection struct {
	connection *gorm.DB
}

func NewUserRepository() *databaseConnection {
	return &databaseConnection{connection: database.ConnectToDB()}
}

func (db *databaseConnection) CloseDB() {
	database.CloseDB(db.connection)
}

func (db *databaseConnection) Save(user entity.User) {
	db.connection.Create(&user)
}

func (db *databaseConnection) Update(user entity.User) {
	db.connection.Save(&user)
}

func (db *databaseConnection) Delete(user entity.User) {
	db.connection.Delete(&user)
}

func (db *databaseConnection) FindAll() []entity.User {
	var users []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&users)
	return users
}
