package configs

import (
	"context"
	"log"

	"github.com/multimoml/extractor/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConfig(key string) (string, error) {
	configServer, err := GetEnv("CONFIG_SERVER")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(*configServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("Connecting to config server at", *configServer)

	if err != nil {
		return "", err
	}
	defer func(conn *grpc.ClientConn) {
		if conn.Close() != nil {
			log.Println("Error closing connection to config server:", err)
		}
	}(conn)

	client := proto.NewConfigClient(conn)
	value, err := client.GetConfig(context.Background(), &proto.ConfigRequest{Key: key})

	if err != nil {
		return "", err
	}

	return value.Value, nil
}
