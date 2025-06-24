FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

COPY ./.env .

RUN go mod download

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api-server

FROM golang:1.24-alpine

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8080

CMD ["./main"]
