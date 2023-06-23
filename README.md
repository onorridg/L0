# L0
![image](https://github.com/onorridg/L0/assets/83474704/a7e836bd-5836-4748-9b66-5b35897c6db5)


.env example:
```.env
# postgresql
PG_HOST=localhost
PG_PORT=5432
PG_USER=admin
PG_PASSWORD=admin
PG_DATABASE=service

# nats-streaming
NATS_PORT=4222
NATS_PORT_HTTP=8222
NATS_CLUSTER_ID=L0
NATS_SUBJECT=order
NATS_GROUP=order-workers
NATS_DURABLE_NAME=order-workers
NATS_PG_DATABASE=order

# worker
WORKER_QUANTITY=10
WORKER_SHUTDOWN_TIME_SECONDS=2

# inMemory (cache)
CACHE_SIZE=100
```
