package gateway

import (
	"database/sql"
	"github.com/tinrab/cautious-giggle/config"
	_ "github.com/lib/pq"
	"context"
)

type DatabaseGateway struct {
	db *sql.DB
}

func NewDatabaseGateway(cfg config.PostgresConfig) (*DatabaseGateway, error) {
	db, err := sql.Open("postgres", cfg.Address)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DatabaseGateway{
		db: db,
	}, nil
}

func (g *DatabaseGateway) Close() {
	g.db.Close()
}

func (g *DatabaseGateway) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := g.db.ExecContext(ctx, query, args...)
	return err
}

func (g *DatabaseGateway) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return g.db.QueryContext(ctx, query, args...)
}
