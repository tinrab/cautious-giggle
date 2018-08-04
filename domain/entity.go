package domain

import (
	"github.com/segmentio/ksuid"
	"time"
)

type Entity struct {
	ID        string
	CreatedAt time.Time
}

func newEntity() Entity {
	return Entity{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
	}
}
