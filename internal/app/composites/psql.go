package composites

import (
	"context"
	"demoApi/internal/app/config"
	"demoApi/pkg/postgresql"
)

type PostgresComposite struct {
	db postgresql.Client
}

func NewPostgresComposite(ctx context.Context, conf config.Config) (*PostgresComposite, error) {
	db, err := postgresql.NewClient(ctx, 5, conf)
	if err != nil {
		return nil, err
	}

	return &PostgresComposite{
		db: db,
	}, nil
}
