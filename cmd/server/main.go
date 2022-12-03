package main

import (
	"context"
	"extractor/internal/server"
)

func main() {
	ctx := context.Background()
	server.Run(ctx)
}
