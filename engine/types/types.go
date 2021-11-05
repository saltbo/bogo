package types

type Routes interface {
	Group(name string) Routes
	NewRs(name Resource)
}

type Router interface {
	Group(name string) Routes
	Run()
}

type Context interface {
}

type Resource interface {
	Get(name string) (interface{}, error)
	Find(query interface{}) (interface{}, error)
	Create(resource interface{}) error
	Update(resource interface{}) error
	Delete(resource interface{}) error
	Query() interface{}
	Model() interface{}
	Name() string
}
