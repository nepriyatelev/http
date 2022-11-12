FROM golang:latest

WORKDIR /app

COPY . .

RUN  go build -o main /app/cmd/main

ENTRYPOINT exec go run cmd/main/main.go