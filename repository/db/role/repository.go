// Package role ...
package role

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

// IRole ...
type IRole interface {
	Create(ctx context.Context, role *model.Role) error
	GetByName(ctx context.Context, name string) (*model.Role, error)
	GetAll(ctx context.Context) ([]model.Role, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]model.Role, error)
}
