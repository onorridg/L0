package main

import (
	"l0/internal/env"
	"l0/internal/server"
	"l0/internal/worker"
	"l0/pkg/inMemory"
)

func main() {
	// init env and inMemory (cache)
	_ = env.Get()
	_ = inMemory.Conn()

	// run worker for nuts streaming and run frontend server
	go worker.Run()
	server.Run()

	//i := inMemory.Node{}
	//node := i.Update(0)
	//for i := 1; i < 10; i++ {
	//	node = node.Append(uint64(i))
	//}
	////fmt.Println("Node length:", node.Len())
	//node = node.FindFirstNode()
	////last := head.FindLastNode()
	//fmt.Println("First:", node.Id, "| Last:", node.Id, "| Node length:", node.Len())
	////update := last.Update(6)
	////fmt.Println("First:", update.FindFirstNode().Id, "| Last:", update.FindLastNode().Id, "| Node length:", last.Len())
	////node = update.Append(10)
	////fmt.Println("First:", node.FindFirstNode().Id, "| Last:", node.FindLastNode().Id, "| Node length:", node.Len())
	////node = update.Update(0)
	////fmt.Println("First:", node.FindFirstNode().Id, "| Last:", node.FindLastNode().Id, "| Node length:", node.Len())
	////node = update.Update(2)
	////fmt.Println("First:", node.FindFirstNode().Id, "| Last:", node.FindLastNode().Id, "| Node length:", node.Len())
	////node = update.Update(0)
	////fmt.Println("First:", node.FindFirstNode().Id, "| Last:", node.FindLastNode().Id, "| Node length:", node.Len())
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	//node = node.Delete()
	////node = node.Delete()
	//
	//fmt.Println("First:", node.FindFirstNode().Id, "| Last:", node.FindLastNode().Id, "| Node length:", node.Len())
}
