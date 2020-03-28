package utils

import (
	"log"

	"github.com/joho/godotenv"
)

//LoadEnv load the env files
func LoadEnv(dir string) bool {
	// loads values from .env into the system
	err := godotenv.Load(dir + ".env")
	if err != nil {
		log.Print("No .env file found")
		return false
	}
	return true
}
