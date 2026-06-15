FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-server ./cmd/api-server
RUN go build -o worker ./cmd/worker

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/api-server .
COPY --from=builder /app/worker .

EXPOSE 8080