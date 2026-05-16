package events

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NATSBus struct {
	conn *nats.Conn
}

func NewNATSBus(url string) (*NATSBus, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATSBus{conn: conn}, nil
}

func (b *NATSBus) Publish(ctx context.Context, topic string, event Event) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return b.conn.Publish(topic, body)
}

func (b *NATSBus) Subscribe(ctx context.Context, topic string, handler Handler) error {
	_, err := b.conn.Subscribe(topic, func(msg *nats.Msg) {
		var event Event

		if err := json.Unmarshal(msg.Data, &event); err != nil {
			return
		}

		_ = handler(ctx, event)
	})

	return err
}

func (b *NATSBus) Close() error {
	b.conn.Close()
	return nil
}
