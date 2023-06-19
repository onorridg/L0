package main

import (
	"github.com/nats-io/stan.go"
	"io"
	"l0/internal/env"
	"log"
	"os"
)

func main() {
	// Подключение к Nats Streaming Server
	cluster := env.Get().NatsClusterId

	sc, err := stan.Connect(cluster, "onorridg")
	if err != nil {
		log.Fatalf("Failed to connect to Nats Streaming: %v", err)
	}
	defer sc.Close()

	//// Отправка сообщения в канал
	//err = sc.Publish("order", []byte("Test group"))
	//if err != nil {
	//	log.Fatalf("Failed to publish message: %v", err)
	//}

	jsonFile, err := os.Open("cmd/sender/model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = sc.Publish("order", byteJson)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Println("Message published successfully.")
}
