package main

import (
	"fmt"
	"l0/internal/env"
)

func main() {
	// Подключение к Nats Streaming Server
	//cluster := env.Get().NatsClusterId
	fmt.Println(env.Get().PgPort, env.Get().NatsPort, env.Get().NatsPortHttp)

	//sc, err := stan.Connect(cluster, "onorridg")
	//if err != nil {
	//	log.Fatalf("Failed to connect to Nats Streaming: %v", err)
	//}
	//defer sc.Close()
	//
	//// Отправка сообщения в канал
	//err = sc.Publish("order", []byte("Test group"))
	//if err != nil {
	//	log.Fatalf("Failed to publish message: %v", err)
	//}
	//
	//log.Println("Message published successfully.")
}
