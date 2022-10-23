package functions

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Env use godot package to load/read the .env file and
//return the value of the key
func Env(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
