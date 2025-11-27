package query

import (
	"context"
	"database/sql"
	"github.com/alihdrii/go-hexagonal/internal/core/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserQueryRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Get(ctx context.Context, conditions map[string]interface{}) (*domain.User, error) {
	return nil, nil
}
