package gateway

import (
	"github.com/nats-io/go-nats"
	"github.com/tinrab/cautious-giggle/config"
)

type EventGateway struct {
	nc *nats.Conn
}

func NewEventGateway(cfg config.NatsConfig) (*EventGateway, error) {
	nc, err := nats.Connect(cfg.Address)
	if err != nil {
		return nil, err
	}
	return &EventGateway{
		nc: nc,
	}, nil
}

func (g *EventGateway) Close() {
	if g.nc != nil {
		g.nc.Close()
	}
}

func (g *EventGateway) Publish(key string, data []byte) error {
	return g.nc.Publish(key, data)
}

func (g *EventGateway) Subscribe(key string, cb func(data []byte)) error {
	_, err := g.nc.Subscribe(key, func(msg *nats.Msg) {
		cb(msg.Data)
	})
	return err
}
