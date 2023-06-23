PHONY: sender services polygon

sender:
	go run cmd/sender/sender.go

services:
	go run cmd/services/services.go

polygon:
	go run cmd/polygon/polygon.go