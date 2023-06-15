package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	// Создание подключения к серверу NATS Streaming
	sc, err := stan.Connect("L0", "worker-2", stan.NatsURL(nats.DefaultURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Определение функции обработчика сообщений
	handler := func(msg *stan.Msg) {
		log.Printf("Получено сообщение: %s\n", string(msg.Data))
		// Обработка сообщения
	}

	// Подписка на канал с использованием опции DurableName и названия группы подписчиков
	sub, err := sc.QueueSubscribe(
		"order",
		"order-workers",
		handler,
		stan.DurableName("order-workers"),
		stan.MaxInflight(10)) // подписчик может обрабатывать не более 10 сообщений одновременно

	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Ожидание завершения программы
	select {}
}
