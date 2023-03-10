package database

import (
	"fmt"
	"os"

	"github.com/DigantaChauduri06/commons"
	"github.com/DigantaChauduri06/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global DB
var DB *gorm.DB

func SetUp() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := os.Getenv("username")
	password := os.Getenv("password")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	commons.CheckError(err)
	db.AutoMigrate(models.Book{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}