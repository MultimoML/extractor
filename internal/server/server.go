package server

import (
	"context"
	"extractor-timer/internal/configs"
	"extractor-timer/internal/controllers"
	"extractor-timer/internal/scraper"
	"fmt"
	"time"
)

func Run(ctx context.Context) {
	start := time.Now()

	// load ENV
	configs.LoadEnvironment()

	products := scraper.ScrapeSpar()

	// run database
	mongoClient := configs.ConnectDB()
	controllers.WriteProductsSpar(ctx, mongoClient, products)

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("Total run time %s\n", elapsed)

	//router := gin.Default()

	// routes
	//routes.Routes(router)

	//router.Run("0.0.0.0:6000")
}
