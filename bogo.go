package bogo

import (
	"github.com/saltbo/bogo/types"
)

type App struct {
	Config   types.ConfigInterface
	Router   types.RouterInterface
	Database types.DatabaseInterface
	Cache    types.CacheInterface
	Logger   types.LoggerInterface
}

func NewApp(config string) (*App, error) {
	app, _, err := buildInjectors(config)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	return nil
}
