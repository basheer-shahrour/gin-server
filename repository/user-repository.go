package repository

import (
	"os"

	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(user entity.User)
	FindAll() []entity.User
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewUserRepository() *database {
	godotenv.Load()
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Conversation{})
	return &database{connection: db}

}

func (db *database) CloseDB() {
	sqlDB, _ := db.connection.DB()
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
}

func (db *database) Save(user entity.User) {
	db.connection.Create(&user)
}

func (db *database) Update(user entity.User) {
	db.connection.Save(&user)
}

func (db *database) Delete(user entity.User) {
	db.connection.Delete(&user)
}

func (db *database) FindAll() []entity.User {
	var users []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&users)
	return users
}
