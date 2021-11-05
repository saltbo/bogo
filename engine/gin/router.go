package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/saltbo/bogo/engine/types"
)

type Router struct {
	*gin.RouterGroup
}

func NewRouter() *Router {
	engine := gin.New()
	return &Router{
		engine.Group(""),
	}
}

func (r *Router) Group(name string) types.Routes {
	return r.engine.Group(name)
}

func (r *Router) NewRs(resource types.Resource) {
	h := &Handler{resource}
	r.engine.Handle("GET", "/"+resource.Name()+"/:name", h.Get)
	r.engine.Handle("GET", "/"+resource.Name(), h.Find)
	r.engine.Handle("POST", "/"+resource.Name(), h.Create)
	r.engine.Handle("PUT", "/"+resource.Name()+"/:name", h.Put)
	r.engine.Handle("DELETE", "/"+resource.Name()+"/:name", h.Delete)
}

func (r *Router) Run() {
	r.engine.Run(":8000")
}
