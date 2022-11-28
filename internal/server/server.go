package server

import (
	"extractor-timer/internal/configs"
	"extractor-timer/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	// load ENV
	configs.LoadEnvironment()

	// run database
	configs.ConnectDB()
	router := gin.Default()

	// routes
	routes.Routes(router)

	router.Run("0.0.0.0:6000")
}
