package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//APIPort port of API
	APIPort = 0
	connStr = ""
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
		log.Print(err)
		log.SetPrefix("WARNING: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		APIPort = 5000
		log.Printf("API port cannot be loaded from environment variables. Defined for %v", APIPort)
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Print(err)
		log.SetPrefix("WARNING: ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		dbPort = 5432
		log.Printf("DB port cannot be loaded from environment variables. Defined for %v", dbPort)
	}

	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName)

}
