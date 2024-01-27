package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var dbConfig *DbConfig

func NewDbConfig() *DbConfig {
	if dbConfig == nil {
		dbConfig = initializeDbConfig()
	}
	return dbConfig
}

func initializeDbConfig() *DbConfig {
	_ = godotenv.Load()
	return &DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
}
