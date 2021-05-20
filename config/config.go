package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBConnectString is the statement to connect the application with the database
	DBConnectString = ""

	// DB_PORT represents the database port
	DB_PORT = 5432

	// API_PORT
	API_PORT = 0
)

// LoadEnvs will initialize the environment variables for the application
func LoadEnvs() {

	// Load .env file
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Get Port and check port availability
	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 8000
	}

	// Get DB related variables
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("POSTGRES_DB")

	DBConnectString = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

}
