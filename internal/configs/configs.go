package configs

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() string {
	environment, err := GetEnv("ENVIRONMENT")
	if err != nil {
		panic(err)
	}

	switch *environment {
	case "dev":
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			fmt.Println("Loaded 'dev' env.")
		}
	case "prod":
		fmt.Println("Loaded 'prod' env.")
	default:
		log.Fatal("Env variable ENVIRONMENT is not set or set to wrong value")
	}

	return *environment
}

func GetEnv(envName string) (*string, error) {
	if envName == "" {
		return nil, errors.New("envName is empty")
	}

	envValue := os.Getenv(envName)

	if envValue == "" {
		return nil, fmt.Errorf("envValue for envName %s is empty", envName)
	}

	return &envValue, nil
}
