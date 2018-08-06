package repositories

import (
	"github.com/tinrab/cautious-giggle/config"
	"github.com/tinrab/cautious-giggle/domain"
	"database/sql"
	"context"
	_ "github.com/lib/pq"
)

type StoreRepository struct {
	db *sql.DB
}

func NewStoreRepository(cfg config.StoreRepositoryConfig) (*StoreRepository, error) {
	db, err := sql.Open("postgres", cfg.Address)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &StoreRepository{
		db,
	}, nil
}

func (r *StoreRepository) Close() {
	r.db.Close()
}

func (r *StoreRepository) InsertProduct(ctx context.Context, p domain.Product) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO products(id, created_at, name, price) VALUES($1, $2, $3, $4)",
		p.ID, p.CreatedAt, p.Name, p.Price,
	)
	return err
}
