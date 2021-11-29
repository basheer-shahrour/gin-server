package database

import (
	"fmt"
	"os"

	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnection interface {
	Close()
}

var Database = NewDatabaseConnection()

func NewDatabaseConnection() *gorm.DB {
	godotenv.Load()
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Conversation{})
	fmt.Printf("connectedToDB\n")
	return db
}

func Close() {
	sqlDB, _ := Database.DB()
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DB CLOSED\n")
}
