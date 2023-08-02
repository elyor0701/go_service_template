package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Global variables
}

func Load() Config {
	var (
		c Config
	)
	if err := godotenv.Load("/app/.env"); err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println("No .env file found")
		}
	} // @TODO

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
