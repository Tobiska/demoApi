package user

import (
	"context"
	"demoApi/internal/app/domain/user"
)

type Service interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*user.User, error)
	Create(ctx context.Context, dto *user.CreateUserDto) error
}
