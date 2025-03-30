package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	log.Print("Load env file success")
}

func GetValue(key string) string {
	return os.Getenv(key)
}
