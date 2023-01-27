package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/app"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/config"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	logger      *log.Logger
	mongoClient *mongo.Client
	aRepo       repository.Account
)

func init() {

	// Load the cofiguration
	cfg := config.LoadConfig()

	// Configure the logger to be used
	logger = log.New(os.Stdout, "decimal-demo-mongo ", log.LstdFlags)

	// Create the Mongo DB Client
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDBConnString))
	if err != nil {
		logger.Fatalf("Error while creating MongoDB client: %#+v\n", err)
	}

	// Create a 10 second time out context & connect to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Fatalf("Error while connecting to MongoDB: %#+v\n", err)
	}

	// Check whether connection is successful or not
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatalf("Error while pinging the MongoDB: %#+v\n", err)
	}
	mongoClient = client
}

func main() {

	// Create a 10 second time out context & connect to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// defer mongoClient connection
	defer mongoClient.Disconnect(ctx)

	// Listing out the database names (Not necessary just doing it for DEMO)
	dbNames, err := mongoClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		logger.Fatalf("Error while retrieving the database names: %#+v", err)
	}

	logger.Println("---------- Current Databases in the MongoDB server ----------")
	logger.Println(dbNames)

	// Get database for our demo (database will be created if it does not exist)
	dbDemo := mongoClient.Database("mongo-demo")
	accountCollection := dbDemo.Collection("accounts")

	// Initizlize the repository to be used
	aRepo = repository.NewAccountImplMongo(logger, accountCollection)

	// Execute the demo application with specified repository
	app.RunDecimalDemo(logger, aRepo)

}
