package user

import (
	"context"
	"demoApi/internal/app/domain/user"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) (*user.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*user.User, error)
	Create(ctx context.Context, dto *CreateUserDto) error
}
