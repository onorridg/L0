PHONY: producer services first_start

producer:
	go run cmd/producer/producer.go

first_start:
	cat .env.example > .env
	docker-compose up

services:
	go run cmd/services/services.go

