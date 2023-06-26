package main

import (
	"l0/internal/env"
	"l0/internal/inMemory"
	"l0/internal/server"
	"l0/internal/worker"
)

func main() {
	// init .env and inMemory (cache)
	_ = env.Get()
	_ = inMemory.Conn()

	// run worker for nuts streaming and run frontend server
	go worker.Run()
	server.Run()
}
