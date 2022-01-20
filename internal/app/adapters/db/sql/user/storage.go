package user

import (
	"context"
	"demoApi/internal/app/domain/user"
	"demoApi/pkg/postgresql"
)

type Storage struct {
	client postgresql.Client
}

func NewStorage(client *postgresql.Client) *Storage {
	return &Storage{
		client: *client,
	}
}

func (st *Storage) GetOne(ctx context.Context, id int) (*user.User, error) {
	u := &user.User{}
	if err := st.client.QueryRow(ctx, `SELECT users.id, users.name, users.email FROM users WHERE users.id=$1`, id).Scan(&u.Id, &u.Name, &u.Email); err != nil {
		return nil, err
	}

	return u, nil
}

func (st *Storage) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	panic("implement GET ALL USERS")
}

func (st *Storage) Create(ctx context.Context, u *user.User) error {
	q := `INSERT INTO users (name, email, password_encrypted) VALUES ($1, $2, $3)`
	if _, err := st.client.Query(ctx, q, u.Name, u.Email, u.EncryptedPassword); err != nil {
		return err
	}
	return nil
}

func (st *Storage) Delete(ctx context.Context, id int) error {
	panic("implement DELETE USER")
}
