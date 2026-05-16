package events

import "context"

type Handler func(context.Context, Event) error

type Bus interface {
	Publish(ctx context.Context, topic string, event Event) error
	Subscribe(ctx context.Context, topic string, handler Handler) error
	Close() error
}
