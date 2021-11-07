package types

import (
	"context"

	"gorm.io/gorm"

	"github.com/saltbo/bogo/router/types"
)

type RouterInterface interface {
	types.Routes
	Run() error
	Shutdown(ctx context.Context) error
}

type DatabaseInterface interface {
	Default() *gorm.DB
	W(name string) *gorm.DB
	R(name string) *gorm.DB
}

type CacheInterface interface {
}

type LoggerInterface interface {
}
