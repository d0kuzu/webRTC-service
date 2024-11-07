package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	ringostatUsername string
	ringostatPassword string
}

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
