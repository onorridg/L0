# L0

.env
```.env
# postgresql
PG_USER=admin
PG_PASSWORD=admin
PG_PORT=5432
PG_DATABASE=service

# nats-streaming
NATS_PORT=4222
NATS_PORT_HTTP=8222
NATS_CLUSTER_ID=L0
NATS_SUBJECT=order
NATS_GROUP=order-workers
NATS_DURABLE_NAME=order-workers
NATS_PG_DATABASE=order
```
