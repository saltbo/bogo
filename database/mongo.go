package database

import (
	"gorm.io/gorm"
)

type Mongo struct {
}

func (m *Mongo) Default() *gorm.DB {
	panic("implement me")
}

func (m *Mongo) W(name string) *gorm.DB {
	panic("implement me")
}

func (m *Mongo) R(name string) *gorm.DB {
	panic("implement me")
}
