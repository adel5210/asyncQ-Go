package main

import (
	"asyncq_test/tasks"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main(){
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default": 3,
				"low": 1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDeliverym, tasks.HandleEmailDeliveryTask)
	mux. Handle(tasks.TypeImageResize, tasks.NewImageProcessor())

	if err:=srv.Run(mux), err != nil{
		log.Fatalf("could not run server: %v", err)
	}
}