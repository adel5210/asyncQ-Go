package main

import (
	"asyncq_test/tasks"
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})
	defer client.Close()

	task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("Could not create task : %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("Could not enqueue task : %v", err)
	}
	log.Printf("Enqueued task: id: %s queue %s", info.ID, info.Queue)
}
