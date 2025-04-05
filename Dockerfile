FROM golang:1.23 AS builder

LABEL authors="FRS"
LABEL version="1.0"
LABEL description="🥷 A RESTful Samurai API built with Go and MongoDB. Tested with honor using Testcontainers."



WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ ./vendor
COPY . .

RUN go build -mod=vendor -o samurai_api ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o samurai_api ./main.go


FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/samurai_api .
COPY internal/banner/ascii.txt ./internal/banner/ascii.txt

EXPOSE 1600

CMD ["./samurai_api"]