package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func initModule() {

	//LoadEnv initially load env

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)

	}

}
