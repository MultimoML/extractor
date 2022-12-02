# https://makefiletutorial.com/

.PHONY: dev build-deploy build deploy

ENVIRONMENT:=dev
export ENVIRONMENT

dev:
	go run cmd/server/main.go

build-deploy:
	make build && \
	make deploy

build:
	sudo docker build -t multimoml/extractor-timer:latest .

deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate