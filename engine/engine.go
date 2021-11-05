package engine

import (
	"github.com/saltbo/bogo/engine/gin"
	"github.com/saltbo/bogo/engine/types"
)

var routers = map[string]types.Router{
	"gin": gin.NewRouter(),
}

func DefaultHandler() types.Router {
	return routers["gin"]
}

type Engine struct {
	types.Router
}

func New() (*Engine, error) {
	return &Engine{
		Router: DefaultHandler(),
	}, nil
}

func (e *Engine) WithRouter(name string) {
	e.Router = routers[name]
}

func (e *Engine) Run() {
	e.Router.Run()
}
