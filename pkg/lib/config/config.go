package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
	Port      string `env:"API_PORT" envDefault:"3000"`
	DBAddress string `env:"API_DB_ADDRESS" envDefault:"mysql:3306"`
	DBName    string `env:"API_DB_NAME" envDefault:"dramas"`
	DBUser    string `env:"API_DB_USER" envDefault:"root"`
}

// Config（環境変数）の読み取り
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("fail to parse cfg: %w", err)
	}
	return cfg, nil
}
