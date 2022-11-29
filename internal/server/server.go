package server

import (
	"extractor-timer/internal/configs"
	"extractor-timer/internal/scraper"
)

func Run() {
	// load ENV
	configs.LoadEnvironment()

	scraper.Scraper()

	// run database
	//configs.ConnectDB()
	//router := gin.Default()

	// routes
	//routes.Routes(router)

	//router.Run("0.0.0.0:6000")
}
