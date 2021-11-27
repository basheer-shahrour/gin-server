package database

import (
	"fmt"
	"os"

	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	ConnectToDB() *gorm.DB
	CloseDB()
}

func ConnectToDB() *gorm.DB {
	godotenv.Load()
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("connectedToDB\n")
	db.AutoMigrate(&entity.User{}, &entity.Conversation{})
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DB CLOSED\n")
}
