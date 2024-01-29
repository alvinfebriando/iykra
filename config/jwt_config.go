package config

import (
	"os"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	Secret string
}

var jwtConfig *JwtConfig

func NewJwtConfig() *JwtConfig {
	if jwtConfig == nil {
		jwtConfig = initializeJwtConfig()
	}
	return jwtConfig
}

func initializeJwtConfig() *JwtConfig {
	_ = godotenv.Load()
	return &JwtConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}
