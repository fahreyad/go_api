package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable is not set")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV environment variable is not set")
	}

	return Config{
		Port: port,
		Env:  env,
	}
}
