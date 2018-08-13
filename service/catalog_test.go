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
	"time"
	"github.com/tinrab/cautious-giggle/listener"
)

var (
	ccs *CatalogCommandService
	cqs *CatalogQueryService
)

func TestMain(m *testing.M) {
	cfg := config.Config{
		Database: config.PostgresConfig{
			Address: "postgres://giggle:123456@localhost:/giggle?sslmode=disable",
		},
		Event: config.NatsConfig{
			Address: "nats://localhost:4222",
		},
	}
	dg, err := gateway.NewDatabaseGateway(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	eg, err := gateway.NewEventGateway(cfg.Event)
	if err != nil {
		log.Fatal(err)
	}

	pr := repository.NewProductRepository(dg)
	ccs = NewCatalogCommandService(eg)
	cqs = NewCatalogQueryService(pr)
	_, err = listener.NewCatalogListener(eg, pr)
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	dg.Close()
	eg.Close()
	os.Exit(code)
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

	<-time.After(100 * time.Millisecond)

	products, err := cqs.GetProductsWithIDs(ctx, ids)
	assert.NoError(t, err)

	for i := 0; i < 3; i++ {
		p := products[i]
		assert.Equal(t, fmt.Sprintf("Product #%d", i+1), p.Name)
		assert.Equal(t, float64(i+1)*2.0, p.Price)
		assert.Equal(t, ids[i], p.ID)
	}
}
