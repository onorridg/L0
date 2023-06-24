PHONY: producer services polygon

producer:
	go run cmd/producer/producer.go

services:
	go run cmd/services/services.go

polygon:
	go run cmd/polygon/polygon.go