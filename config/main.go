package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config app config data
type Config struct {
	LineChannel string `env:"LINE_CHANNEL"`
}

// GetConfig will return app configuration from
// environment variable and config file
func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}
	err = env.Parse(&config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return config
}
