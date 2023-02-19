package main

import (
	"github.com/DigantaChauduri06/commons"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	commons.CheckError(err)
}