// Package kafka ...
package kafka

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	"context"

	"github.com/lukmanlukmin/go-lib/kafka"
)

// IRepository ...
type IRepository interface {
	Publish(ctx context.Context, topic string, value interface{}) error
}

// Repository ...
type Repository struct {
	Producer kafka.Producer
}

// NewRepository ...
func NewRepository(producer kafka.Producer) *Repository {
	return &Repository{
		Producer: producer,
	}
}
