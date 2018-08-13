package service

import (
	"github.com/tinrab/cautious-giggle/domain"
	"context"
	"github.com/tinrab/cautious-giggle/gateway"
	"encoding/gob"
	"bytes"
)

type CatalogCommandService struct {
	eventGateway *gateway.EventGateway
}

func NewCatalogCommandService(eg *gateway.EventGateway) *CatalogCommandService {
	return &CatalogCommandService{
		eventGateway: eg,
	}
}

func (s *CatalogCommandService) InsertProduct(ctx context.Context, p domain.Product) error {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(p)
	if err != nil {
		return err
	}
	return s.eventGateway.Publish("catalog:product:create", b.Bytes())
}
