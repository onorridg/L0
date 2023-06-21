package main

import (
	"l0/internal/env"
	"l0/internal/server"
	"l0/internal/worker"
)

func main() {
	_ = env.Get()
	go worker.Run()
	server.Run()
}
