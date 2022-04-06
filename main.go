package main

import (
	"log"

	"insanitygaming.net/bans/services/database"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvVariables()
	database.Connect()
	defer database.Disconnect()

}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
