package db

import (
	"context"
	"github.com/alihdrii/go-hexagonal/internal/core/domain"
)

type UserCommandRepository interface {
	Store(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
}

type UserQueryRepository interface {
	Get(ctx context.Context, conditions map[string]interface{}) (*domain.User, error)
}
