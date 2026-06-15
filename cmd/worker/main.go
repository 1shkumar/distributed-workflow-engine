package main

import (
	"log"

	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/kafka"
	"github.com/1shkumar/mini-orchestrator/internal/worker"
)

func main() {

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Worker service started")

	kafka.InitProducer()

	worker.StartConsumer()
}
