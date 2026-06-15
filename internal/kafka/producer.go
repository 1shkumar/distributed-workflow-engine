package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
)

var Writer *kafkago.Writer

func InitProducer() {

	Writer = &kafkago.Writer{
		Addr: kafkago.TCP(
			"orchestrator-kafka:9092",
		),
		Topic:    "task-queue",
		Balancer: &kafkago.LeastBytes{},
	}

	fmt.Println("Kafka producer initialized")
}

func PublishTask(
	payload interface{},
) error {

	message, err :=
		json.Marshal(payload)

	if err != nil {
		return err
	}

	return Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: message,
		},
	)
}
