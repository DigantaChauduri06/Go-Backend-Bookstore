package database

import (
	"github.com/DigantaChauduri06/models"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil,err
	}
	return books, nil
}

func GetBookById(db *gorm.DB, id string) models.Book {
	return models.Book{}
}

func DeleteBook(db *gorm.DB, id string) error {
	return nil
}

func UpdateBook(db *gorm.DB, book *models.Book) error {
	return nil
}