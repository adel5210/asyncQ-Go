package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceURL string
}

func NewEmailDeliveryTask(usrID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{
		UserID:     usrID,
		TemplateID: tmplID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{
		SourceURL: src,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("Json unmarshal failed : %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to user: user_id: %d, template_id: %s", p.UserID, p.TemplateID)
	return nil
}

type ImageProcessor struct {
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("Json unmarshal failed : %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resize image: src: %s", p.SourceURL)
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}
