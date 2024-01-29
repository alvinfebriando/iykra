package config

import (
	"os"

	"github.com/joho/godotenv"
)

type RestConfig struct {
	Host string
	Port string
}

var restConfig *RestConfig

func NewRestConfig() *RestConfig {
	if restConfig == nil {
		restConfig = initializeRestConfig()
	}
	return restConfig
}

func initializeRestConfig() *RestConfig {
	_ = godotenv.Load()

	host := os.Getenv("REST_HOST")
	port := os.Getenv("REST_PORT")

	return &RestConfig{
		Host: host,
		Port: port,
	}
}
