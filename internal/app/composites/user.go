package composites

import (
	user3 "demoApi/internal/app/adapters/db/sql/user"
	"demoApi/internal/app/adapters/handlers"
	user2 "demoApi/internal/app/adapters/handlers/user"
	"demoApi/internal/app/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user2.Service
	Handler handlers.Handler
}

func NewUserComposite(composite *PostgreSQLComposite) (*UserComposite, error) {
	st := user3.NewStorage(&composite.db)
	sv := user.NewService(st)
	h := user2.NewHandler(sv)

	return &UserComposite{
		Storage: st,
		Service: sv,
		Handler: h,
	}, nil
}
