// Package userrole ...
package userrole

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=mocks

import (
	model "attendance/model/db"
	"context"

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

// IUserRole ...
type IUserRole interface {
	Create(ctx context.Context, userRole *model.UserRole) error
}
