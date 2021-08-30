package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port   int    `env:"PORT" envDefault:"3000"`
	DBConn string `env:"DB_CONN" envDefault:"mongodb://localhost"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
