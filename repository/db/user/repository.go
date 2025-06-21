// Package user ...
package user

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

// IUser ...
type IUser interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, ID uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, uname string) (*model.User, error)
}
