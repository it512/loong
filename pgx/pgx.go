package pgx

import (
	"context"
	"database/sql"

	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent"
	_ "github.com/it512/loong/pgx/internal/ent/runtime"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	*ent.Client
}

func open(db *sql.DB, ops ...ent.Option) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(append(ops, ent.Driver(drv))...)
}

func OpenDBCtx(ctx context.Context, url string) (*sql.DB, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err

	}
	return stdlib.OpenDBFromPool(pool), nil
}

func OpenDB(url string) (*sql.DB, error) {
	return OpenDBCtx(context.Background(), url)
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Client: open(db),
	}
}

func PgxStore(url string) loong.Option {
	db := loong.Must(OpenDB(url))
	store := NewStore(db)
	return loong.SetStore(store)
}
