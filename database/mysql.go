package database

import (
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/saltbo/bogo/config"
)

type FieldID struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type FieldTimes struct {
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type MySQL struct {
	dbs map[string]*gorm.DB
}

func New(c *config.Config) (*MySQL, error) {
	dbs := make(map[string]*gorm.DB, 0)
	for key, database := range c.Database {
		var director func(dsn string) gorm.Dialector
		switch database.Type {
		case "mysql":
			director = mysql.Open
		case "postgres":
			director = postgres.Open
		case "sqlserver":
			director = sqlserver.Open
		case "clickhouse":
			director = clickhouse.Open
		default:
			director = sqlite.Open
		}

		db, err := gorm.Open(director(database.DSN), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		// todo 连接池配置化
		sqlDB, err := db.DB()
		sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。

		if c.Logger.Level == "debug" {
			db = db.Debug()
		}

		dbs[key] = db
	}

	return &MySQL{
		dbs: dbs,
	}, nil
}

func (d *MySQL) Default() *gorm.DB {
	if db, ok := d.dbs["default"]; ok {
		return db
	}

	return nil
}

// W WriteDB
func (d *MySQL) W(name string) *gorm.DB {
	if db := d.Select(name + "_write"); db != nil {
		return db
	}

	return d.Select(name)
}

// R ReadDB
func (d *MySQL) R(name string) *gorm.DB {
	if db := d.Select(name + "_read"); db != nil {
		return db
	}

	return d.Select(name)
}
func (d *MySQL) Select(name string) *gorm.DB {
	if db, ok := d.dbs[name]; ok {
		return db
	}

	return nil
}
