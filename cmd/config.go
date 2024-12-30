package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PAT string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	return &Config{
		PAT: os.Getenv("PAT"),
	}, nil
}
