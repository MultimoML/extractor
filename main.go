package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/multimoml/extractor/internal/server"
)

// @title Extractor API
// @version 1.0.0
// @host localhost:6000
// @BasePath /extractor
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go server.Run(ctx)

	<-sigChan
	cancel()
}
