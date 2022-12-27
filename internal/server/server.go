package server

import (
	"context"
	"extractor/internal/configs"
	"extractor/internal/db_client"
	"extractor/internal/internal_state"
	"extractor/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

const serviceName = "extractor"

func Run(ctx context.Context) {
	// load ENV
	environment := configs.LoadEnvironment()

	// Logging
	tracer.Start(
		tracer.WithService(serviceName),
		tracer.WithEnv(environment),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService(serviceName),
		profiler.WithEnv(environment),
		profiler.WithVersion("0.0.1"),
		profiler.WithTags("router:GinGonic,database:MongoDB"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	// init database
	db_client.DBClient(ctx)

	// init internalState
	internal_state.InternalState(ctx)

	// start http server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(gintrace.Middleware(serviceName))

	routes.Routes(router)
	address := fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))
	router.Run(address)
}
