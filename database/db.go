package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

// Global DB
var DB *gorm.DB

func SetUp() {
	
	
	host := "localhost:5432"
	port := "5432"
	dbName := "book"
	username := os.Getenv("username")
	password := os.Getenv("password")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	log.Println(dsn)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// commons.CheckError(err)
	// db.AutoMigrate(models.Book{})
	// DB = db
}

func GetDB() *gorm.DB {
	return DB
}