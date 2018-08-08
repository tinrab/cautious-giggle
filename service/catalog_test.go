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
)

var (
	ccs *CatalogCommandService
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

	os.Exit(m.Run())
}

func TestInsertProduct(t *testing.T) {
	ctx := context.Background()
	p := domain.NewProduct("A", 1)
	ccs.InsertProduct(ctx, p)
}
