package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GrpcServerAddress string
	DBConnString      string
}

func LoadConfig() (config Config) {

	// Get the gRPC port number
	err := godotenv.Load("app.env")
	if err != nil {
		panic(err)
	}

	// Get the POSTGESS DB variables
	dbName := os.Getenv("POSTGRES_DB")
	if len(dbName) <= 0 {
		panic("POSTGRES_DB IS NOT SET ! ! !")
	}

	dbHost := os.Getenv("POSTGRES_DB_HOST")
	if len(dbHost) <= 0 {
		panic("POSTGRES_DB_HOST IS NOT SET ! ! !")
	}

	dbPort := os.Getenv("POSTGRESS_DB_PORT")
	if len(dbPort) <= 0 {
		panic("POSTGRESS_DB_PORT IS NOT SET ! ! !")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if len(dbUser) <= 0 {
		panic("POSTGRES_USER IS NOT SET ! ! !")
	}

	dbPass := os.Getenv("POSTGRES_PASSWORD")
	if len(dbPass) <= 0 {
		panic("POSTGRES_PASSWORD IS NOT SET ! ! !")
	}

	//config.DBConnString = "host=:: port=5432 user=postgres password=postgres dbname=testAssignment sslmode=disable"
	config.DBConnString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	return
}
