package scraper

import (
	"encoding/json"
	"extractor-timer/internal/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const sparScrapingURL = "https://search-spar.spar-ics.com/fact-finder/rest/v4/search/products_lmos_si?query=*&hitsPerPage=9999999"

func ScrapeSpar() models.Products {
	start := time.Now()

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().Get(sparScrapingURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Spar returned reponse in %.2fs.\n", resp.Time().Seconds())

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

	// fmt.Printf("\n%+v\n", dataUnparsed)

	dataParsed, err := ParseSpar(dataUnparsed, timestamp)
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

	// fmt.Printf("\n%#v\n", products)

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("ScrapeSpar run time %s\n", elapsed)

	return products
}
