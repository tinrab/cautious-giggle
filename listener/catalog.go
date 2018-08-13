package listener

import (
	"github.com/tinrab/cautious-giggle/gateway"
	"github.com/tinrab/cautious-giggle/repository"
	"bytes"
	"github.com/tinrab/cautious-giggle/domain"
	"encoding/gob"
	"log"
	"context"
)

type CatalogListener struct {
	eventGateway      *gateway.EventGateway
	productRepository *repository.ProductRepository
}

func NewCatalogListener(eg *gateway.EventGateway, pr *repository.ProductRepository) (*CatalogListener, error) {
	l := &CatalogListener{
		eventGateway:      eg,
		productRepository: pr,
	}
	err := eg.Subscribe("catalog:product:create", l.OnProductCreate)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *CatalogListener) OnProductCreate(data []byte) {
	b := bytes.Buffer{}
	b.Write(data)
	p := domain.Product{}
	err := gob.NewDecoder(&b).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	err = l.productRepository.CreateProduct(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
}
