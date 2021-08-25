package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port         int    `env:"PORT" envDefault:"3000"`
	DBConnString string `env:"DB_CONN_STRING" envDefault:"mongodb://127.0.0.1"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
