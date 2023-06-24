# L0
![image](https://github.com/onorridg/L0/assets/83474704/a7e836bd-5836-4748-9b66-5b35897c6db5)


## Подготовка окружения:
Создать в корне проекта файл `.env`:
```.env
# postgresql
PG_HOST=localhost
PG_PORT=5432
PG_USER_ADMIN=admin
PG_PASSWORD_ADMIN=admin
PG_DATABASE=service
PG_USER_WORKER=worker
PG_PASSWORD_WORKER=worker

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
Поднятие postgresql и nats streaming:
```bash
docker-compoe up -d
```
Подготовка базы данных (зависимость от psql):
```bash
PGPASSWORD=admin psql -h localhost -p 5432 -U admin -d service  -f database/postgres/models.sql
PGPASSWORD=admin psql -h localhost -p 5432 -U admin -d service  -f database/postgres/worker.sql
```
Запуск сервиса (worker, inMemory и frontend server)
```bash
make service
```
