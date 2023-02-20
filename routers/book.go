package routers

import (
	"github.com/DigantaChauduri06/database"
	"github.com/DigantaChauduri06/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	api := handlers.Handlers{
		DB: database.GetDB(),
	}

	router.GET("/get-books", api.GetBooks)
	return router
}