package user

import (
	"context"
	"demoApi/internal/app/adapters/handlers/user"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) user.Service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	panic("implement me!!!")
}
func (s *service) GetAll(ctx context.Context, limit, offset int) ([]*User, error) {
	panic("implement me!!!")
}
func (s *service) Create(ctx context.Context, dto *user.CreateUserDto) error {
	panic("implement me!!!")
}
