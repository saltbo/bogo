package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Logger struct {
	Level string `yaml:"level" toml:"level"`
	Path  string `yaml:"path" toml:"path"`
}

type Database struct {
	Type  string `yaml:"type" toml:"type"`
	DSN   string `yaml:"dsn" toml:"dsn"`
}

type Config struct {
	Env      string              `yaml:"env" toml:"env"`
	Port     int                 `yaml:"port" toml:"port"`
	Logger   Logger              `yaml:"logger" toml:"logger"`
	Database map[string]Database `yaml:"database" toml:"database"`
}

func New(filepath string) (*Config, error) {
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.AddConfigPath(filepath)                // path to look for the config file in
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	var cfg Config
	return &cfg, viper.Unmarshal(&cfg)
}

func (c *Config) DefaultDB() Database {
	for _, wr := range c.Database {
		return wr
	}

	return Database{}
}

func (c *Config) SelectDB(name string) Database {
	return c.Database[name]
}
