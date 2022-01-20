package user

import (
	"context"
	"demoApi/internal/app/domain/user"
	"errors"
)

type Storage struct {
	users map[int]*user.User
}

func NewStorage(users *map[int]*user.User) *Storage {
	return &Storage{
		users: *users,
	}
}

func (st *Storage) GetOne(ctx context.Context, id int) (*user.User, error) {
	u, ke := st.users[id]
	if ke == false {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (st *Storage) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	var um []*user.User
	if len(st.users) > 0 {
		um = make([]*user.User, 0, offset+limit)
		for i := offset; i < (offset + limit); i++ {
			if _, ke := st.users[i]; ke != false {
				//	um[i] = st.users[i]
				um = append(um, st.users[i])
			}
		}
	}
	return um, nil
}

func (st *Storage) Create(ctx context.Context, u *user.User) error {
	st.users[u.Id] = u
	return nil
}

func (st *Storage) Delete(ctx context.Context, id int) error {
	lm := len(st.users)
	delete(st.users, id)
	//cause delete no-op
	if lm == len(st.users) {
		return errors.New("no user by that id")
	}
	return nil
}
