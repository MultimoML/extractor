package main

import (
	"context"
	"extractor/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go server.Run(ctx)

	<-sigChan
	cancel()
}
