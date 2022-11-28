package scraper

import (
	"encoding/json"
	"extractor-timer/internal/models"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const sparScrapingURL = "https://search-spar.spar-ics.com/fact-finder/rest/v4/search/products_lmos_si?query=*&hitsPerPage=99999999"

func Scraper() {
	var products models.Products

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().Get("sparScrapingURL")
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Returned non 200 status code %v.", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &products)
	if err != nil {
		log.Fatal(err)
	}
}
