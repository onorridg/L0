PHONY: producer services polygon init_env

producer:
	go run cmd/producer/producer.go

services:
	go run cmd/services/services.go

polygon:
	go run cmd/polygon/polygon.go



init_env:
	docker-compose up -d
	@sleep 20
	PGPASSWORD=admin psql -h localhost -p 5432 -U admin -d service  -f database/postgres/models.sql
	PGPASSWORD=admin psql -h localhost -p 5432 -U admin -d service  -f database/postgres/worker.sql
	service
	echo "[+] Создано рабочее окружение. Сервисы запущены."
	echo "[+] Frontend доступен на http://localhost:8080"