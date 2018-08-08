package service

import (
	"github.com/tinrab/cautious-giggle/domain"
	"github.com/tinrab/cautious-giggle/repository"
	"context"
)

type CatalogCommandService struct {
	productRepository *repository.ProductRepository
}

func NewCatalogCommandService(pr *repository.ProductRepository) *CatalogCommandService {
	return &CatalogCommandService{
		productRepository: pr,
	}
}

func (s *CatalogCommandService) InsertProduct(ctx context.Context, p domain.Product) (error) {
	err := s.productRepository.CreateProduct(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
