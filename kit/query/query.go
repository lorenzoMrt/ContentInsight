package query

import (
	"context"
)

//go:generate mockery --case=snake --outpkg=querymocks --output=querymocks --name=Bus

type Bus interface {
	Register(Type, Handler)
	Ask(context.Context, Query) (interface{}, error)
}

type Type string

type Query interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Query) (interface{}, error)
}
