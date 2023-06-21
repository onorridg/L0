package worker

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"l0/internal/env"
	"l0/internal/models"
	"l0/internal/postgresql"
	"l0/pkg/convert"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type workerData struct {
	ctx context.Context
	sc  stan.Conn
	sub stan.Subscription
	id  uint8
	db  *postgresql.DB
}

func (w *workerData) msgHandler(msg *stan.Msg) {
	order := models.Order{}
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		log.Println("[!] msgHandler:", err)
		return
	}
	w.db.InsertUserOrder(&order)
}

func worker(wD *workerData) {
	defer func() {
		if err := wD.sub.Unsubscribe(); err != nil {
			log.Println("sub:", wD.id, err)
		}
		if err := wD.sc.Close(); err != nil {
			log.Println("sc:", wD.id, err)
		}
		if err := wD.db.Conn.Close(); err != nil {
			log.Println("db close:", wD.id, err)
		}
	}()

	select {
	case <-wD.ctx.Done():
		//fmt.Println("[+] Done goroutine:", wD.id)
	}
}

func runWorkers(ctx context.Context) {
	wQ := env.Get().WorkerQuantity
	for ; wQ > 0; wQ-- {
		wD := workerData{ctx: ctx, id: uint8(wQ - 1)}
		idStr := "worker-" + convert.NumToStr(wD.id)

		var err error
		if wD.sc, err = stan.Connect("L0", idStr, stan.NatsURL(nats.DefaultURL)); err != nil {
			log.Fatal(err)
		}
		if wD.sub, err = wD.sc.QueueSubscribe("order", "order-workers", wD.msgHandler,
			stan.DurableName("order-workers"), stan.MaxInflight(10)); err != nil { // подписчик может обрабатывать не более 10 сообщений одновременно
			log.Fatal(err)
		}
		wD.db = postgresql.Conn()
		go worker(&wD)
	}
	log.Printf("[+] %d workers launched", env.Get().WorkerQuantity)
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	runWorkers(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	log.Println("[!] Shutdown workers ...")

	shutdownTime := env.Get().WorkerShutdownTime
	time.Sleep(time.Second * shutdownTime)

	log.Println("[+] Workers exiting")
}
