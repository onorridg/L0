package worker

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func msgHandler(msg *stan.Msg) {
	log.Printf("Получено сообщение: %s\n", string(msg.Data))
}

type workerData struct {
	ctx context.Context
	sc  stan.Conn
	sub stan.Subscription
	id  uint8
}

func worker(wD *workerData) {
	defer func() {
		err := wD.sub.Unsubscribe()
		if err != nil {
			log.Println("sub:", wD.id, err)
		}
		err = wD.sc.Close()
		if err != nil {
			log.Println("sc:", wD.id, err)
		}
	}()

main:
	for {
		select {
		case <-wD.ctx.Done():
			fmt.Println("Закругляется:", wD.id)
			break main
		default:
			fmt.Println("Работает:", wD.id)
			time.Sleep(time.Second * 3)
		}
	}

}

func runWorkers(ctx context.Context) {
	for i := 0; i < 10; i++ {
		wD := workerData{}
		wD.ctx = ctx
		wD.id = uint8(i)
		id := "worker-" + strconv.Itoa(i)
		var err error
		wD.sc, err = stan.Connect("L0", id, stan.NatsURL(nats.DefaultURL))
		if err != nil {
			log.Fatal(err)
		}
		//defer sc.Close()
		wD.sub, err = wD.sc.QueueSubscribe("order", "order-workers", msgHandler,
			stan.DurableName("order-workers"),
			stan.MaxInflight(10)) // подписчик может обрабатывать не более 10 сообщений одновременно
		if err != nil {
			log.Fatal(err)
		}
		go worker(&wD)
		//defer sub.Unsubscribe()
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	runWorkers(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	log.Println("Shutdown workers ...")
	time.Sleep(time.Second * 8)
	log.Println("Workers exiting")
}
