package db_client

import (
	"context"
	"fmt"
	"log"
	"os"
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

	mUsername := os.Getenv("M_USERNAME")
	mPassword := os.Getenv("M_PASSWORD")

	mongodbUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/", mUsername, mPassword, "localhost", "27017")

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

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	database := os.Getenv("DATABASE")

	collection := client.Database(database).Collection(collectionName)
	return collection
}
