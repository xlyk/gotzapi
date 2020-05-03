.phony: build run login push

build:
	docker-compose build

run:
	docker-compose up -d

push:
	docker-compose push

login:
	cat ./token.txt | docker login docker.pkg.github.com -u xlyk --password-stdin

