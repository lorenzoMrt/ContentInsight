package inmemory

import (
	"context"
	"log"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/query"
)

// CommandBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Dispatch implements the query.Bus interface.
func (b *QueryBus) Ask(ctx context.Context, cmd query.Query) (interface{}, error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return cr.ContentResponse{}, nil
	}

	r, err := handler.Handle(ctx, cmd)
	if err != nil {
		log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
	}
	resp := r.(cr.ContentResponse)

	return resp, nil
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(cmdType query.Type, handler query.Handler) {
	b.handlers[cmdType] = handler
}
