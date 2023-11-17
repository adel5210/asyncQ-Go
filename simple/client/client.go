package main

import (
	"asyncq_test/simple/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr: redisAddr,
		},
	)
	defer client.Close()

	immediateEnqueue(client)
	scheduledQueue(client)
}

func immediateEnqueue(client *asynq.Client) {
	// Immediate enqueue
	task, err := tasks.NewTopic0Task("admin")
	if err != nil {
		log.Fatal("Cannot create task on topic 0")
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatal("Cannot enqueue task on topic 0")
	}

	log.Printf("Topic 0 task is on queued with id: %s, queue: %s", info.ID, info.Queue)

}

func scheduledQueue(client *asynq.Client) {
	task, err := tasks.NewTopic0Task("admin")
	if err != nil {
		log.Fatal("Cannot create task on topic 0")
	}

	info, err := client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal("Cannot enqueue task on topic 0")
	}

	log.Printf("Topic 0 task is on queued with id: %s, queue: %s", info.ID, info.Queue)

}
