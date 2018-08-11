package service

import (
	"github.com/tinrab/cautious-giggle/repository"
	"context"
	"github.com/tinrab/cautious-giggle/domain"
)

type CatalogQueryService struct {
	productRepository *repository.ProductRepository
}

func NewCatalogQueryService(pr *repository.ProductRepository) *CatalogQueryService {
	return &CatalogQueryService{
		productRepository: pr,
	}
}

func (s *CatalogQueryService) GetProductsWithIDs(ctx context.Context, ids []string) ([]domain.Product, error) {
	return s.productRepository.ReadProductsWithIDs(ctx, ids)
}
