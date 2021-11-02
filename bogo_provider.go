package bogo

import (
	"github.com/saltbo/bogo/config"
	"github.com/saltbo/bogo/router"
	"github.com/saltbo/bogo/types"
)

func ConfigConstructor(filepath string) (types.ConfigInterface, error) {
	return config.NewDefaultConfig()
}

func RouterConstructor(configInterface types.ConfigInterface) (types.RouterInterface, error) {
	return router.NewDefaultRouter()
}

func DatabaseConstructor(configInterface types.ConfigInterface) (types.DatabaseInterface, error) {
	return nil, nil
}

func CacheConstructor(configInterface types.ConfigInterface) (types.CacheInterface, error) {
	return nil, nil
}

func LoggerConstructor(configInterface types.ConfigInterface) (types.LoggerInterface, error) {
	return nil, nil
}
