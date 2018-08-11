package service

import (
	"testing"
	"os"
	"github.com/tinrab/cautious-giggle/config"
	"github.com/tinrab/cautious-giggle/repository"
	"github.com/tinrab/cautious-giggle/gateway"
	"log"
	"context"
	"github.com/tinrab/cautious-giggle/domain"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var (
	ccs *CatalogCommandService
	cqs *CatalogQueryService
)

func TestMain(m *testing.M) {
	cfg := config.Config{
		Database: config.DatabaseConfig{
			Address: "postgres://giggle:123456@localhost:/giggle?sslmode=disable",
		},
	}
	dg, err := gateway.NewDatabaseGateway(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	pr := repository.NewProductRepository(dg)
	ccs = NewCatalogCommandService(pr)
	cqs = NewCatalogQueryService(pr)

	os.Exit(m.Run())
}

func TestInsertAndGetProducts(t *testing.T) {
	ctx := context.Background()

	var ids []string

	for i := 1; i <= 3; i++ {
		p := domain.NewProduct(fmt.Sprintf("Product #%d", i), float64(i)*2.0)
		err := ccs.InsertProduct(ctx, p)
		assert.NoError(t, err)
		ids = append(ids, p.ID)
	}

	products, err := cqs.GetProductsWithIDs(ctx, ids)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		p := products[i]
		assert.Equal(t, fmt.Sprintf("Product #%d", i+1), p.Name)
		assert.Equal(t, float64(i+1)*2.0, p.Price)
		assert.Equal(t, ids[i], p.ID)
	}
}
