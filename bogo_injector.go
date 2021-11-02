//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package bogo

import (
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(App), "*"))

func buildInjectors(config string) (*App, func(), error) {
	wire.Build(
		ConfigConstructor,
		RouterConstructor,
		DatabaseConstructor,
		CacheConstructor,
		LoggerConstructor,
		InjectorSet,
	)
	return new(App), nil, nil
}
