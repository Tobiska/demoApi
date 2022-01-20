package user

import "context"

type Storage interface {
	GetOne(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
