package main

import (
	"context"
	"extractor-timer/internal/server"
)

func main() {
	ctx := context.Background()
	server.Run(ctx)
}
