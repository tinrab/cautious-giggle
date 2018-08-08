package repository

import (
	"github.com/tinrab/cautious-giggle/domain"
	"context"
	"github.com/lib/pq"
	"github.com/tinrab/cautious-giggle/gateway"
)

type ProductRepository struct {
	databaseGateway *gateway.DatabaseGateway
}

func NewProductRepository(dg *gateway.DatabaseGateway) *ProductRepository {
	return &ProductRepository{
		databaseGateway: dg,
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p domain.Product) error {
	err := r.databaseGateway.Exec(
		ctx,
		"INSERT INTO products(id, created_at, name, price) VALUES($1, $2, $3, $4)",
		p.ID, p.CreatedAt, p.Name, p.Price,
	)
	return err
}

func (r *ProductRepository) ReadProductsWithIDs(ctx context.Context, ids []string) ([]domain.Product, error) {
	rows, err := r.databaseGateway.Query(
		ctx,
		"SELECT id, created_at, name, price FROM products WHERE id = ANY($1::CHAR(27)[])",
		pq.Array(ids),
	)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	product := &domain.Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.CreatedAt, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}
