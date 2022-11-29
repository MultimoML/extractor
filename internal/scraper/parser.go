package scraper

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/itchyny/gojq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	jqQuerySpar = `.hits[] | 
			{
				id: .id,
				name: .masterValues.name,

				"is-on-promotion": 				.masterValues."is-on-promotion",
				"category-names": 				.masterValues."category-names",
				"category-name": 					.masterValues."category-name",
				"allergens-filter": 			.masterValues."allergens-filter",

				"sales-unit": .masterValues."sales-unit",
				title: .masterValues.title,
				"code-internal": .masterValues."code-internal",

				price: 										.masterValues.price,
				"created-at": 						.masterValues."created-at",
				"best-price": 						.masterValues."best-price",
				"stock-status": 					.masterValues."stock-status",
				"is-new": 								.masterValues."is-new",
				"image-url": 							.masterValues."image-url",
				"approx-weight-product": 	.masterValues."approx-weight-product",
				url: 											.masterValues.url,
				brand: 										.masterValues."ecr-brand",
				"price-per-unit": 				.masterValues."price-per-unit",
				"regular-price": 					.masterValues."regular-price",
				"price-per-unit-number": 	.masterValues."price-per-unit-number",
			}
			| .id |= leadingZeros24 
			| ."is-on-promotion" |= tobool
			| ."category-names" |= split
			| ."created-at" |= todatetime
			| ."code-internal" |= tonumber
			| ."is-new" |= tobool
			| ."approx-weight-product" |= tobool
		`
	timeout = time.Minute
)

func ParseSpar(dataUnparsed interface{}) ([]interface{}, error) {
	fmt.Println("Started ParseSpar...")

	query, err := gojq.Parse(jqQuerySpar)
	if err != nil {
		log.Fatal(err)
	}

	code, err := gojq.Compile(query,
		gojq.WithFunction("leadingZeros24", 0, 0, func(x interface{}, xs []interface{}) (out interface{}) {
			number, err := strconv.Atoi(x.(string))
			if err != nil {
				log.Fatal(err)
			}

			return fmt.Sprintf("%024d", number)
		}),
		gojq.WithFunction("tobool", 0, 0, func(x interface{}, xs []interface{}) (out interface{}) {
			boolean, err := strconv.ParseBool(x.(string))
			if err != nil {
				log.Fatal(err)
			}

			return boolean
		}),
		gojq.WithFunction("todatetime", 0, 0, func(x interface{}, xs []interface{}) (out interface{}) {
			datetime, err := strconv.ParseInt(x.(string), 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			return primitive.DateTime(datetime)
		}),
		gojq.WithFunction("split", 0, 0, func(x interface{}, xs []interface{}) (out interface{}) {
			if x == nil {
				return []string{"VSI IZDELKI", "OSTALO"}
			}

			stringRepresentation := x.(string)

			array := strings.Split(stringRepresentation, "|")

			return array
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	var dataParsed []interface{}

	ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
	defer ctxCancel()

	fmt.Println("Starting gojq engine...")
	iter := code.RunWithContext(ctx, dataUnparsed)
	for {
		v, ok := iter.Next()
		// fmt.Printf("\n%+v\n", v)

		if !ok {
			break
		}

		if err, ok := v.(error); ok {
			log.Fatal(err)
		}

		dataParsed = append(dataParsed, v)
	}

	fmt.Println("ParseSpar parser successfully finished parsing.")
	fmt.Printf("Parsed %v entrys.\n", len(dataParsed))

	return dataParsed, nil
}
