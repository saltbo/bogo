package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewDefaultRouter() (*Router, error) {
	engine := gin.New()

	return &Router{
		engine,
	}, nil
}
