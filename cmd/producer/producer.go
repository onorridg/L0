package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"io"
	"l0/internal/env"
	"l0/internal/models"
	"log"
	"os"
)

func main() {
	cluster := env.Get().NatsClusterId
	sc, err := stan.Connect(cluster, "onorridg")
	if err != nil {
		log.Fatalf("Failed to connect to Nats Streaming: %v", err)
	}
	defer sc.Close()

	jsonFile, err := os.Open("cmd/producer/model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var order models.Order
	if err = json.Unmarshal(byteJson, &order); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		order.TrackNumber = uuid.New().String()
		order.Items[0].TrackNumber = order.TrackNumber
		order.Items[1].TrackNumber = order.TrackNumber
		bOrder, err := json.Marshal(order)
		if err != nil {
			log.Fatal(err)
		}
		err = sc.Publish(env.Get().NatsSubject, bOrder)
		if err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
	}
	log.Println("Message published successfully.")
}
