package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client instance
var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	mUsername := os.Getenv("M_USERNAME")
	mPassword := os.Getenv("M_PASSWORD")

	mongodbUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/", mUsername, mPassword, "localhost", "27017")

	fmt.Println(mongodbUrl)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	database := os.Getenv("DATABASE")

	collection := client.Database(database).Collection(collectionName)
	return collection
}
