FROM golang:1.20-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags "-w -s" -o services ./cmd/services/services.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/services .
COPY .env .
COPY frontend frontend
EXPOSE "${FRONTEND_PORT}" "${FRONTEND_PORT}"
CMD [ "./services" ]