package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	D DatabaseConfig
}

func InitConfig() Config {
	// Load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Print(".env successfully loaded")

	return Config{
		D: MakeDBConfig(),
	}

}

func GetENV(key string, value string) string {
	// Set env by the value if empty
	if ok := os.Getenv(key) == ""; ok {
		os.Setenv(key, value)
	}
	return os.Getenv(key)
}
