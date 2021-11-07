package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/saltbo/bogo/config"
	"github.com/saltbo/bogo/router/types"
)

type Router struct {
	rGroup

	srv    *http.Server
	engine *gin.Engine
}

func NewRouter(cfg *config.Config) *Router {
	engine := gin.Default()
	return &Router{
		rGroup: rGroup{rg: &engine.RouterGroup},

		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: engine,
		},
		engine: engine,
	}
}

func (r *Router) Run() error {
	log.Printf("[rest server listen at %s]", r.srv.Addr)
	if err := r.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (r *Router) Shutdown(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

type rGroup struct {
	rg *gin.RouterGroup
}

func (r *rGroup) Group(name string) types.Routes {
	return &rGroup{
		rg: r.rg.Group(name),
	}
}
func (r *rGroup) NewRs(name string, resource types.Resource) {
	rp1 := fmt.Sprintf("/%s/:name", name)
	rp2 := fmt.Sprintf("/%s", name)
	h := &Handler{resource}
	r.rg.GET(rp1, h.Get)
	r.rg.GET(rp2, h.Find)
	r.rg.POST(rp2, h.Create)
	r.rg.PUT(rp1, h.Put)
	r.rg.DELETE(rp1, h.Delete)
}
