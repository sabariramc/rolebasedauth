package test

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Print("Testing no .env file found")
	}
}
