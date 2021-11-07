package bogo

import (
	"github.com/saltbo/bogo/config"
	"github.com/saltbo/bogo/database"
	"github.com/saltbo/bogo/router"
	"github.com/saltbo/bogo/types"
)

func ConfigConstructor(filepath string) (*config.Config, error) {
	return config.New(filepath)
}

func RouterConstructor(cfg *config.Config) (types.RouterInterface, error) {
	return router.NewRouter(cfg), nil
}

func DatabaseConstructor(cfg *config.Config) (types.DatabaseInterface, error) {
	return database.New(cfg)
}

func CacheConstructor(cfg *config.Config) (types.CacheInterface, error) {
	return nil, nil
}

func LoggerConstructor(cfg *config.Config) (types.LoggerInterface, error) {
	return nil, nil
}
