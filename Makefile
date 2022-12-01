# https://makefiletutorial.com/

.PHONY: build deploy

ENVIRONMENT:=dev
export ENVIRONMENT

run:
	go run cmd/server/main.go

build-deploy:
	make build && \
	make deploy

build:
	sudo docker build -t extractor-timer .

deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate