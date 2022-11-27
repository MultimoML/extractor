# https://makefiletutorial.com/

.PHONY: build deploy

build:


deploy:
	sudo docker-compose -f docker-compose.yml --env-file .env up -d --force-recreate