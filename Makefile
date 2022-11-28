# https://makefiletutorial.com/

.PHONY: build deploy

run:
	export ENVIRONMENT=dev
	go run main.go

build:


deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate