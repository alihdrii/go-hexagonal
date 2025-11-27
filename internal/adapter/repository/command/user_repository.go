package command

import (
	"context"
	"database/sql"
	"github.com/alihdrii/go-hexagonal/internal/core/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserCommandRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Store(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}
