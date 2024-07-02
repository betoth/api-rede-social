package config

import (
	"api-rede-social/src/logs"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//APIPort port of API
	APIPort = 0

	//ConnStr Connection String
	ConnStr = ""

	SecretKey []byte
)

// LoadEnv Loading all enviroment variables
func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 5000
		logs.Warning(err, fmt.Sprintf("API port cannot be loaded from environment variables. Defined for %v", APIPort))

	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		dbPort = 5432
		logs.Warning(err, fmt.Sprintf("DB port cannot be loaded from environment variables. Defined for %v", dbPort))

	}

	ConnStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
