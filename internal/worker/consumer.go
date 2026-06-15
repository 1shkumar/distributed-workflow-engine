package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	kafkago "github.com/segmentio/kafka-go"

	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/kafka"
	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func StartConsumer() {

	reader := kafkago.NewReader(
		kafkago.ReaderConfig{
			Brokers: []string{
				"orchestrator-kafka:9092",
			},
			Topic:       "task-queue",
			GroupID:     "worker-group",
			StartOffset: kafkago.FirstOffset,
			MinBytes:    10e3,
			MaxBytes:    10e6,
		},
	)

	log.Println("Kafka consumer started")

	for {

		log.Println(
			"Waiting for Kafka message...",
		)

		message, err :=
			reader.ReadMessage(
				context.Background(),
			)

		if err != nil {
			log.Println(
				"failed to consume message:",
				err,
			)
			continue
		}

		log.Println(
			"Consumed message:",
			string(message.Value),
		)

		var event models.TaskEvent

		err = json.Unmarshal(
			message.Value,
			&event,
		)

		if err != nil {
			log.Println(
				"failed to parse event:",
				err,
			)
			continue
		}

		log.Println(
			"Executing task:",
			event.TaskName,
		)

		time.Sleep(
			5 * time.Second,
		)

		err = db.MarkTaskSuccess(
			event.WorkflowRunID,
			event.TaskName,
		)

		if err != nil {
			log.Println(
				"failed updating task:",
				err,
			)
			continue
		}

		log.Println(
			"Task completed:",
			event.TaskName,
		)

		workflow, err :=
			db.GetWorkflowDefinitionByRunID(
				event.WorkflowRunID,
			)

		if err != nil {

			log.Println(
				"failed fetching workflow:",
				err,
			)

			continue
		}

		log.Println(
			"Workflow loaded:",
			workflow.Name,
		)

		for _, task := range workflow.Tasks {

			dependentOnCurrent := false

			for _, dependency := range task.DependsOn {

				if dependency ==
					event.TaskName {

					dependentOnCurrent = true
					break
				}
			}

			if !dependentOnCurrent {
				continue
			}

			canRun := true

			for _, dependency := range task.DependsOn {

				if !db.IsTaskCompleted(
					event.WorkflowRunID,
					dependency,
				) {

					canRun = false
					break
				}
			}

			if !canRun {
				continue
			}

			err = db.MarkTaskQueued(
				event.WorkflowRunID,
				task.Name,
			)

			if err != nil {

				log.Println(
					"failed queueing task:",
					err,
				)

				continue
			}

			err = kafka.PublishTask(
				models.TaskEvent{
					WorkflowRunID: event.WorkflowRunID,
					TaskName:      task.Name,
				},
			)

			if err != nil {

				log.Println(
					"failed publishing task:",
					err,
				)

				continue
			}

			log.Println(
				"Queued next task:",
				task.Name,
			)
		}

		if db.AreAllTasksCompleted(
			event.WorkflowRunID,
		) {

			err = db.MarkWorkflowCompleted(
				event.WorkflowRunID,
			)

			if err != nil {

				log.Println(
					"failed marking workflow complete:",
					err,
				)

				continue
			}

			log.Println(
				"Workflow completed:",
				event.WorkflowRunID,
			)
		}
	}
}
