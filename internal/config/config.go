package config

import "time"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxLifetime     time.Duration
	MaxIdleLifetime time.Duration
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Address:      ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  time.Minute,
		},
		Database: DatabaseConfig{
			URL:             "postgres://postgres:postgres@localhost:5432/products?sslmode=disable",
			MaxOpenConns:    25,
			MaxIdleConns:    25,
			MaxLifetime:     5 * time.Minute,
			MaxIdleLifetime: 2 * time.Minute,
		},
	}, nil
}
