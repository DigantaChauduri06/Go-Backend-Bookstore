package main

import (
	"log"

	"github.com/DigantaChauduri06/commons"
	"github.com/DigantaChauduri06/database"
	"github.com/DigantaChauduri06/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	commons.CheckError(err)
	database.SetUp()
	engine := routers.Router()
	if err := engine.Run("localhost:8080"); err != nil {
		log.Fatal("Something unexpected happen while engine start")
	}
}