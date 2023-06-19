package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type data struct {
	PgHost     string
	PgUser     string
	PgPassword string
	PgPort     string
	PgDatabase string

	NatsPort        string
	NatsPortHttp    string
	NatsClusterId   string
	NatsSubject     string
	NatsGroup       string
	NatsDurableName string
	NatsPgDatabase  string

	WorkerQuantity     uint
	WorkerShutdownTime time.Duration
}

var envData *data

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envData = &data{}
	envData.PgHost = os.Getenv("PG_HOST")
	envData.PgUser = os.Getenv("PG_USER")
	envData.PgPassword = os.Getenv("PG_PASSWORD")
	envData.PgPort = os.Getenv("PG_PORT")
	envData.PgDatabase = os.Getenv("PG_DATABASE")

	envData.NatsPort = os.Getenv("NATS_PORT")
	envData.NatsPortHttp = os.Getenv("NATS_PORT_HTTP")
	envData.NatsClusterId = os.Getenv("NATS_CLUSTER_ID")
	envData.NatsSubject = os.Getenv("NATS_SUBJECT")
	envData.NatsGroup = os.Getenv("NATS_GROUP")
	envData.NatsDurableName = os.Getenv("NATS_DURABLE_NAME")
	envData.NatsPgDatabase = os.Getenv("NATS_PG_DATABASE")

	if _, err = fmt.Sscan(os.Getenv("WORKER_QUANTITY"), &envData.WorkerQuantity); err != nil {
		log.Fatal(err)
	}
	if _, err = fmt.Sscan(os.Getenv("WORKER_SHUTDOWN_TIME"), &envData.WorkerShutdownTime); err != nil {
		log.Fatal(err)
	}
}

func Get() *data {
	if envData == nil {
		initEnv()
	}
	return envData
}
