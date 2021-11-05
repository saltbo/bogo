package model

type Meta struct {
	Name string `json:"name"`
}

// Book abc
type Book struct {
	meta Meta `json:"meta"`

	Id   int    `json:"id"`
	Name string `json:"string"`
}
