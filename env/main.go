package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("Error loading .env file\n%v", err)
	}
}
