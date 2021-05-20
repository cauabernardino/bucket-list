package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// DBConnectString is the statement to connect the application with the database
	DBConnectString = ""
	DB_PORT         = 5432
)

// LoadEnvs will initialize the environment variables for the application
func LoadEnvs() {

	// Load .env file
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Get DB related variables
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME := os.Getenv("POSTGRES_DB")

	DBConnectString = fmt.Sprintf(
		"postgres://%s:%s@localhost:%d/%s?sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_PORT,
		DB_NAME,
	)

}
