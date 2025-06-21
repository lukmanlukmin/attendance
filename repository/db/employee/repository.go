// Package employee ...
package employee

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	model "attendance/model/db"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

// IEmployee ...
type IEmployee interface {
	Create(ctx context.Context, emp *model.Employee) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.Employee, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Employee, error)
	GetBatch(ctx context.Context, batchSize, offset int) ([]model.Employee, error)
}
