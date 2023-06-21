package main

import (
	"l0/internal/env"
	"l0/internal/server"
)

func main() {
	_ = env.Get()
	//worker.Run()
	server.Run()
}
