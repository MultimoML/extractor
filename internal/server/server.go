package server

import (
	"context"
	"extractor-timer/internal/configs"
	"extractor-timer/internal/db_client"
	"extractor-timer/internal/internal_state"
	"extractor-timer/internal/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) {
	// load ENV
	configs.LoadEnvironment()

	// init database
	db_client.DBClient(ctx)

	// init internalState
	internal_state.InternalState(ctx)

	// start http server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.Routes(router)
	address := fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))
	router.Run(address)
}
