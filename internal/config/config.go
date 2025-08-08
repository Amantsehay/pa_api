package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DatabaseDSN string
    // Add more config fields as needed
}

var AppConfig *Config

func LoadConfig() {
    _ = godotenv.Load()
    AppConfig = &Config{
        DatabaseDSN: os.Getenv("DATABASE_DSN"),
    }
    if AppConfig.DatabaseDSN == "" {
        log.Fatal("DATABASE_DSN is required but not set")
    }
}