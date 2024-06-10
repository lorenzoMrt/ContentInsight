package inmemory

import (
	"context"

	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	events []event.Event
}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (b *EventBus) Publish(_ context.Context, events []event.Event) error {
	b.events = append(b.events, events...)
	return nil
}
