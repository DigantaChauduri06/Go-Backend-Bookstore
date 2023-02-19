package models

import "gorm.io/gorm"


type Book struct {
	gorm.Model
	Author string `json:"author"`
	Price int `json:"price"`
	Pages int `json:"page"`
	BookName string `json:"bookname"`
	Id string `json:"id"`
}
