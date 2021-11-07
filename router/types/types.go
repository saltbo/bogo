package types

import "context"

type Routes interface {
	Group(name string) Routes
	NewRs(name string, resource Resource)
}

type Context interface {
}

type Query interface{}

type Resource interface {
	Query() Query
	Model() interface{}
	Get(ctx context.Context, name string) (interface{}, error)
	Find(ctx context.Context, query Query) (interface{}, int64, error)
	Create(ctx context.Context, resource interface{}) error
	Update(ctx context.Context, resource interface{}) error
	Delete(ctx context.Context, resource interface{}) error
}
