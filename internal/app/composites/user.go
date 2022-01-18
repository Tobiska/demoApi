package composites

import (
	"demoApi/internal/app/adapters/handlers"
	user2 "demoApi/internal/app/adapters/handlers/user"
	"demoApi/internal/app/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user2.Service
	Handler handlers.Handler
}

func NewUserComposite() (*UserComposite, error) {
	// TODO add postgresql
	return &UserComposite{}, nil
}
