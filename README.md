# L0
![image](https://github.com/onorridg/L0/assets/83474704/a7e836bd-5836-4748-9b66-5b35897c6db5)


## Первый запуск сервиса(L0) и окружения:
```bash
cat .env.example > .env
docker-compose up
```
## Повторный запуск/остановка сервиса и окружения:
```bash
docker-compose start  # запуск
docker-compose stop   # остановка
```
---
## Запуск продюсера:
```bash
make producer
```
Продюсер отправит в `nats streaming` 5 сообщений.

--- 
## Общая информация:
- `Producer`: отвечает за отправку сообщений в `nats streaming`.
- `Services`:
    - `worker`: отвечает за обработку сообщений из `nats streaming`, записывает данные в `postgresql`.
    - `frontend server`: 
        - Endpoints:
            - http://localhost:8080/ отвечает за интерфейс.
            - http://localhost:8080/get-json?id=1 отвечает за выдачу `.json` по порядковому `id` из базы.
        - Отвечает за добавление данных в `inMemory`. Данные добавляются в кэш, если были вызваны хотя бы один раз на стороне пользователя.
        - Удаление данных из `inMemory`. После каждого добавления в кэш, если `.env:CACHE_SIZE` == `len(inMemory)`, то удаляется запись из кэша, которая не использовалась дольше всего.
        - Если данные были выданы из кэша, то в заголовке `X-Cache-Status` будет `Hit`. Если данные были получены из `postgresql`, то заголовок будет равен `Miss`.       

Чтобы узнать порядковый `id`, для обращения к нему через интерфейс, можно заглянуть в логи `services`. Так же в логах можно увидеть удачное или неудачное добавление сообщений в `postgresql`.

У nats streaming в качестве хранилища установлен `postgresql`, так что, при выключении 'nats streaming', необработанные сообщения никуда не пропадут.
