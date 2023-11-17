package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const (
	TypeTopic0 = "username:register"
)

type Topic0Payload struct {
	Id    uint32
	Token string
}

func NewTopic0Task(id uint32, token string) (*asynq.Task, error) {
	payload, err := json.Marshal(
		Topic0Payload{
			Id:    id,
			Token: token,
		},
	)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeTopic0, payload), nil
}

func HandleTopic0Task(ctx context.Context, t *asynq.Task) error {
	var p Topic0Payload
	if err := json.Unmarshal(
		t.Payload(),
		&p,
	); err != nil {
		return fmt.Errorf("Json unmarshal failed on topic 0, cause: %s \n", err.Error())
	}

	log.Printf("Processing topic 0 payload %+v...", p)
	return nil
}
