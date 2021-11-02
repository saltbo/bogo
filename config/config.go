package config

type Config struct {
}

func NewDefaultConfig() (*Config, error) {
	return &Config{}, nil
}

func (c *Config) Get(key string) {
	panic("implement me")
}
