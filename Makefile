# https://makefiletutorial.com/

.PHONY: build deploy

run:
	export ENVIRONMENT=dev
	go run internal/server/main.go

build:
	go build internal/server/main.go

deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate