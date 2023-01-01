package db_client

import (
	"context"
	"extractor/internal/configs"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once sync.Once

	ctx      context.Context
	dbClient *mongo.Client
)

func DBClient(ctxIn ...context.Context) *mongo.Client {

	once.Do(func() {
		ctx = ctxIn[0]
		connectDB()
	})

	return dbClient
}

func connectDB() {
	fmt.Printf("\nStarted ConnectDB...\n")

	dbUsername, err := configs.GetEnv("DB_USERNAME")
	if err != nil {
		panic(err)
	}
	dbPassword, err := configs.GetEnv("DB_PASSWORD")
	if err != nil {
		panic(err)
	}
	dbHost, err := configs.GetEnv("DB_HOST")
	if err != nil {
		panic(err)
	}

	mongodbUrl := fmt.Sprintf("mongodb://%s:%s@%s/", *dbUsername, *dbPassword, *dbHost)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, ctxCancel := context.WithTimeout(ctx, 10*time.Second)
	defer ctxCancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbClient = client
	fmt.Printf("Connected to MongoDB\n")
}

func GetCollectionInternalState(client *mongo.Client) *mongo.Collection {
	databaseName, err := configs.GetEnv("DATABASE_INTERNAL_STATE")
	if err != nil {
		panic(err)
	}

	collectionName, err := configs.GetEnv("COLLECTION_INTERNAL_STATE")
	if err != nil {
		panic(err)
	}

	collection := getCollection(client, *databaseName, *collectionName)
	return collection
}

func GetCollectionExtractor(client *mongo.Client, collectionName string) *mongo.Collection {
	databaseName, err := configs.GetEnv("DATABASE_EXTRACTOR")
	if err != nil {
		panic(err)
	}

	collection := getCollection(client, *databaseName, collectionName)
	return collection
}

// getting database collections
func getCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
