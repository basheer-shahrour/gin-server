package database

import (
	"os"

	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	ConnectToDB() *database
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func (db *database) ConnectToDB() {
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
