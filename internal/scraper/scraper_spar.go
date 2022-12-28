package scraper

import (
	"context"
	"encoding/json"
	"extractor/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const sparScrapingURL = "https://search-spar.spar-ics.com/fact-finder/rest/v4/search/products_lmos_si?query=*&hitsPerPage=9999999"

func ScrapeSpar(ctx context.Context) models.Products {
	fmt.Printf("\nStarted ScrapeSpar...\n")
	start := time.Now()

	// Create a Resty Client
	client := resty.New().SetTimeout(time.Minute)

	resp, err := client.R().Get(sparScrapingURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Spar returned reponse in %.2fs\n", resp.Time().Seconds())

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Returned non 200 status code %v.", resp.StatusCode())
	}

	rawData := resp.Body()
	timestamp := resp.ReceivedAt()

	var dataUnparsed interface{}

	err = json.Unmarshal(rawData, &dataUnparsed)
	if err != nil {
		log.Fatal(err)
	}

	dataParsed, err := parseSpar(ctx, dataUnparsed, timestamp)
	if err != nil {
		log.Fatal(err)
	}

	dataEncoded, err := json.Marshal(dataParsed)
	if err != nil {
		log.Fatal(err)
	}

	var products models.Products

	err = json.Unmarshal(dataEncoded, &products)
	if err != nil {
		log.Fatal(err)
	}

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("ScrapeSpar run time %v\n", elapsed)

	return products
}
