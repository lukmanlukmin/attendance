// Package server ...
package server

import (
	"attendance/bootstrap"
	"attendance/config"
	"attendance/constant"
	"attendance/handler/event"
	"context"

	kafkalib "github.com/lukmanlukmin/go-lib/kafka"
	"github.com/lukmanlukmin/go-lib/log"
)

// ConsumerRouter ...
func ConsumerRouter(ctx context.Context, b *bootstrap.Bootstrap, cfg *config.Config) {
	consumer := kafkalib.NewConsumerGroup(&cfg.KafkaConfig)

	handler := event.NewHandler(b)
	consumer.Subscribe(&kafkalib.ConsumerContext{
		Context: ctx,
		Topics:  []string{constant.TopicCalculatePayroll},
		GroupID: cfg.KafkaConfig.ClientID,
		Handler: kafkalib.MessageProcessorFunc(func(msg *kafkalib.MessageDecoder) {
			err := handler.CalculatePayroll(ctx, msg.Body)
			if err != nil {
				msg.Commit(msg)
			} else {
				log.WithContext(ctx).WithError(err).Error("failed to calculate payroll")
			}
		}),
	})
}
