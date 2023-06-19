PHONY: run sender worker polygon

sender:
	go run cmd/sender/sender.go

worker:
	go run cmd/worker/worker.go

run:
	go run cmd/server/main.go

polygon:
	go run cmd/polygon/polygon.go