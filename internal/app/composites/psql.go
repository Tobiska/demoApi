package composites

import (
	"context"
	"demoApi/internal/app/config"
	"demoApi/pkg/postgresql"
)

type PostgreSQLComposite struct {
	db postgresql.Client
}

func NewPostgreSQLComposite(ctx context.Context, conf config.Config) (*PostgreSQLComposite, error) {
	db, err := postgresql.NewClient(ctx, 5, conf)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLComposite{
		db: db,
	}, nil
}
