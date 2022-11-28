package main

import (
	"extractor-timer/configs"
	"extractor-timer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// load ENV
	configs.LoadEnvironment()

	// run database
	configs.ConnectDB()
	router := gin.Default()

	// routes
	routes.UserRoute(router)

	router.Run("0.0.0.0:6000")
}
