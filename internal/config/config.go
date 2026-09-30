package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return &Config{}, fmt.Errorf("Error loading env")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if(dbUrl == "") {
		return &Config{}, fmt.Errorf("DB URL is required")
	}

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	
	return &Config{
		DatabaseURL: dbUrl,
		Port: port,
	}, nil
}