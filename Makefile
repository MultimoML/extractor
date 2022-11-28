# https://makefiletutorial.com/

.PHONY: build deploy

ENVIRONMENT:=dev
export ENVIRONMENT

run:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go

deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate