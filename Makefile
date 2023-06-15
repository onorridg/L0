PHONY: run sender worker

sender:
	go run cmd/sender/sender.go

worker:
	go run cmd/worker/worker.go

run:
	go run cmd/server/main.go

