package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type data struct {
	PgUser     string
	PgPassword string
	PgPort     uint16
	PgDatabase string

	NatsPort        uint16
	NatsPortHttp    uint16
	NatsClusterId   string
	NatsSubject     string
	NatsGroup       string
	NatsDurableName string
	NatsPgDatabase  string
}

var envData *data

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envData = &data{}
	envData.PgUser = os.Getenv("PG_USER")
	envData.PgPassword = os.Getenv("PG_PASSWORD")
	fmt.Sscan(os.Getenv("PG_PORT"), &envData.PgPort)

	fmt.Sscan(os.Getenv("NATS_PORT"), &envData.NatsPort)
	fmt.Sscan(os.Getenv("NATS_PORT_HTTP"), &envData.NatsPortHttp)
	envData.NatsClusterId = os.Getenv("NATS_CLUSTER_ID")
	envData.NatsSubject = os.Getenv("NATS_SUBJECT")
	envData.NatsGroup = os.Getenv("NATS_GROUP")
	envData.NatsDurableName = os.Getenv("NATS_DURABLE_NAME")
	envData.NatsPgDatabase = os.Getenv("NATS_PG_DATABASE")
}

func Get() *data {
	if envData == nil {
		initEnv()
	}
	return envData
}
