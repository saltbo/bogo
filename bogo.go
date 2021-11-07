package bogo

import (
	"context"
	"log"
	"time"

	"github.com/saltbo/bogo/config"
	"github.com/saltbo/bogo/pkg/cmd"
	"github.com/saltbo/bogo/types"
)

type App struct {
	Config   *config.Config
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
	var err error
	go func() {
		err = a.Router.Run()
	}()
	
	cmd.CatchSignals(a.Stop)
	return err
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Router.Shutdown(ctx); err != nil {
		log.Fatal("[app shutdown err:]", err)
	}

	select {
	case <-ctx.Done():
		log.Println("[app exit timeout of 5 seconds.]")
	default:

	}

	log.Printf("[app exited.]")
}
