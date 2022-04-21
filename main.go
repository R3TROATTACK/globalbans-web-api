package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"insanitygaming.net/bans/services/database"
	"insanitygaming.net/bans/services/logger"
)

func main() {
	setup := flag.Bool("setup", false, "Install the database")
	flag.Parse()
	var log = logger.Logger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()
	defer database.Close()

	if *setup {
		database.RunSetup()
		return
	}

	background := context.Background()
	rows, err := database.Query(background, "SELECT * FROM gb_bans")
	fmt.Print(rows, err)
}
