package main

import (
	"asyncq_test/simple/tasks"
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redisAddr,
		},
		asynq.Config{
			Concurrency: 4,
			Queues: map[string]int{
				"critical": 3,
				"default":  2,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeTopic0, tasks.HandleTopic0Task)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run AsynQ server %v", err.Error())
	}
}
