package config

import (
	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	DBHost     			string `env:"DB_HOST" required:"false"`
	DBPort           int    `env:"DB_PORT" required:"false"`
	DBUser           string `env:"DB_USER" required:"false"`
	DBPasswd         string `env:"DB_PASSWD" required:"false"`
	DBName           string `env:"DB_NAME" required:"false"`
}

func NewConfig(logger *zap.Logger) (Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("could not parse .env file")
		return Config{}, err
	}

	config := Config{}
	if err := env.Set(&config); err != nil {
		logger.Error("could not parse environment variables")
		return Config{}, err
	}

	return config, err
}
