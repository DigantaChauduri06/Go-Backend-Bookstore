package handlers

import (
	"fmt"
	"net/http"

	"github.com/DigantaChauduri06/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type Handlers struct {
	DB *gorm.DB
}

func (h *Handlers) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		fmt.Println("Error happend for fetching books")
		in.JSON(http.StatusInternalServerError,err)
	} else {
		in.JSON(http.StatusOK,books)
	}


}