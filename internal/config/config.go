package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port string
	JwtSecret string
	AccessTokenTTL time.Duration
	RefreshTokenTTL time.Duration
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return &Config{}, fmt.Errorf("jwt_secret is required")
	}

	accessMinutes := 15
	if accessTokenTTL := os.Getenv("ACCESS_TOKEN_TTL_MINUTES"); accessTokenTTL != "" {
		if n, err := strconv.Atoi(accessTokenTTL); err == nil {
			accessMinutes = n
		}
	}
	
	refreshDays := 30
	if refreshTokenTTL := os.Getenv("REFRESH_TOKEN_TTL_DAYS"); refreshTokenTTL != "" {
		if n, err := strconv.Atoi(refreshTokenTTL); err == nil {
			refreshDays = n
		}
	}
	
	return &Config{
		DatabaseURL: dbUrl,
		Port: port,
		JwtSecret: jwtSecret,
		AccessTokenTTL: time.Duration(accessMinutes) * time.Minute,
		RefreshTokenTTL: time.Duration(refreshDays) * 24 * time.Hour,
	}, nil
}