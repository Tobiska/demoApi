package user

import (
	"context"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetByID(ctx context.Context, id int) (*User, error) {
	u, err := s.storage.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil

}
func (s *service) GetAll(ctx context.Context, limit, offset int) ([]*User, error) {
	us, err := s.storage.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return us, nil
}
func (s *service) Create(ctx context.Context, dto *CreateUserDto) error {
	u := &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	//todo ctx handler timeout or deadline
	if err := u.EncryptPassword(); err != nil {
		return err
	}

	if err := s.storage.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.storage.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
