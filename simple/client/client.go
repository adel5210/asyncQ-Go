package main

import (
	"asyncq_test/simple/tasks"
	"log"
	"time"

	"github.com/google/uuid"
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
	var i uint32
	for i = 0; i < 1000; i++ {
		immediateEnqueue(i, client)
	}
	// scheduledQueue(client)
}

func immediateEnqueue(id uint32, client *asynq.Client) {
	// Immediate enqueue
	uuid := uuid.NewString()
	task, err := tasks.NewTopic0Task(id, uuid)
	if err != nil {
		log.Fatalln("Cannot create task on topic 0 " + uuid)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalln("Cannot enqueue task on topic 0")
	}

	log.Printf("Topic 0 task is on queued with id: %d info.id: %s, queue: %s\n", id, info.ID, info.Queue)

}

func scheduledQueue(client *asynq.Client) {
	task, err := tasks.NewTopic0Task(0, "admin")
	if err != nil {
		log.Fatalln("Cannot create task on topic 0")
	}

	info, err := client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatalln("Cannot enqueue task on topic 0")
	}

	log.Printf("Topic 0 task is on queued with id: %s, queue: %s\n", info.ID, info.Queue)

}
