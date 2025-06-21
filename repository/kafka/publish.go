// Package kafka ...
package kafka

import (
	"attendance/model/event"
	"context"

	"github.com/lukmanlukmin/go-lib/kafka"
)

// Publish ...
func (r *Repository) Publish(ctx context.Context, topic string, value interface{}) error {
	data, err := event.BuildKafkaPayload(value, topic)
	if err != nil {
		return err
	}
	return r.Producer.Publish(ctx, &kafka.MessageContext{
		Topic: topic,
		Value: data,
	})
}
