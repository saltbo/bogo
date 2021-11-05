package main

import (
	"log"

	"github.com/saltbo/bogo"
)

func main() {
	s, err := bogo.NewApp("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	rg := s.Router.Group("api")
	rg.Group("v1").NewRs(&bookHook{})

	// 声明一个ResourceStruct，可以为这个资源设置一些钩子，在钩子里实现业务逻辑
	// 每个资源自动具备CURD逻辑，每个操作触发一些钩子

	s.Run()
}


// hooks.go
type BookHooks struct {
}

type bookHook struct {
}

func (hook *bookHook) Get(name string) (interface{}, error) {
	panic("implement me")
}

func (hook *bookHook) Find(query interface{}) (interface{}, error) {
	panic("implement me")
}

func (hook *bookHook) Create(resource interface{}) error {
	panic("implement me")
}

func (hook *bookHook) Update(resource interface{}) error {
	panic("implement me")
}

func (hook *bookHook) Delete(resource interface{}) error {
	panic("implement me")
}

func (hook *bookHook) Query() interface{} {
	panic("implement me")
}

func (hook *bookHook) Model() interface{} {
	panic("implement me")
}

func (hook *bookHook) Name() string {
	panic("implement me")
}