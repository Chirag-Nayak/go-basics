package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/app"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/config"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	aRepo  repository.Account
	logger *log.Logger
)

func init() {

	// Load the cofiguration
	cfg := config.LoadConfig()

	// Configure the logger to be used
	logger = log.New(os.Stdout, "decimal-demo ", log.LstdFlags)

	// Open connection to the DB
	dbClient, err := sql.Open("postgres", cfg.DBConnString)
	if err != nil {
		panic(err)
	}
	// Check the ping to make sure connection to the DB is successful
	err = dbClient.Ping()
	if err != nil {
		panic(err)
	}
	logger.Println("Database connection is successful !")
	db = dbClient
}

func main() {

	// Initizlize the repository to be used
	aRepo = repository.NewAccountImplPgsql(logger, db)

	// Execute the demo application with specified repository
	app.RunDecimalDemo(logger, aRepo)
}
