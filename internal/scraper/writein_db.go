package scraper

import (
	"context"
	"extractor/internal/db_client"
	"extractor/internal/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WriteProductsSpar(ctx context.Context, products models.Products) {
	fmt.Printf("\nStarting WriteProductsSpar...\n")
	start := time.Now()

	dbClient := db_client.DBClient()
	sparCollection := db_client.GetCollection(dbClient, "spar")

	for _, product := range products {
		resultInsert, err := sparCollection.InsertOne(ctx, product)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				fmt.Println(resultInsert)
				log.Fatal(err)
			}

			filter := bson.M{
				"_id": product.Id,
			}
			update := bson.M{
				"$push": bson.M{
					"price-in-time": product.PriceInTime[0],
				},
			}

			resultUpdate, err := sparCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(resultUpdate)
				log.Fatal(err)
			}
		}
	}

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("WriteProductsSpar run time %v\n", elapsed)
}
