package main

import (
	"context"
	"log"

	"github.com/saltbo/bogo"
	"github.com/saltbo/bogo/router/types"
)

func main() {
	s, err := bogo.NewApp("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	rg := s.Router.Group("api")
	rg.Group("v1").NewRs("books", &bookHook{})

	// 声明一个ResourceStruct，可以为这个资源设置一些钩子，在钩子里实现业务逻辑
	// 每个资源自动具备CURD逻辑，每个操作触发一些钩子

	s.Run()
}

type bookHook struct {
}

func (b bookHook) Query() types.Query {
	panic("implement me")
}

func (b bookHook) Model() interface{} {
	panic("implement me")
}

func (b bookHook) Get(ctx context.Context, name string) (interface{}, error) {
	panic("implement me")
}

func (b bookHook) Find(ctx context.Context, query types.Query) (interface{}, int64, error) {
	panic("implement me")
}

func (b bookHook) Create(ctx context.Context, resource interface{}) error {
	panic("implement me")
}

func (b bookHook) Update(ctx context.Context, resource interface{}) error {
	panic("implement me")
}

func (b bookHook) Delete(ctx context.Context, resource interface{}) error {
	panic("implement me")
}

