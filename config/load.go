package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// function for  load config file
func LoadConfig(key string) string {

	// load .env file from that path
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error load file")
	}

	return os.Getenv(key)
}
