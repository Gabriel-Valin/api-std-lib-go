package config

import "time"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Address string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	URL string

	MaxOpenConns int
	MaxIdleConns int

	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

func Load() (*Config, error) {
	readTimeout, err := getEnvDuration(
		"SERVER_READ_TIMEOUT",
		10*time.Second,
	)
	if err != nil {
		return nil, err
	}

	writeTimeout, err := getEnvDuration(
		"SERVER_WRITE_TIMEOUT",
		10*time.Second,
	)
	if err != nil {
		return nil, err
	}

	idleTimeout, err := getEnvDuration(
		"SERVER_IDLE_TIMEOUT",
		time.Minute,
	)
	if err != nil {
		return nil, err
	}

	maxOpenConns, err := getEnvInt(
		"DB_MAX_OPEN_CONNS",
		25,
	)
	if err != nil {
		return nil, err
	}

	maxIdleConns, err := getEnvInt(
		"DB_MAX_IDLE_CONNS",
		25,
	)
	if err != nil {
		return nil, err
	}

	connMaxLifetime, err := getEnvDuration(
		"DB_CONN_MAX_LIFETIME",
		5*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	connMaxIdleTime, err := getEnvDuration(
		"DB_CONN_MAX_IDLE_TIME",
		2*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Address: getEnv(
				"SERVER_ADDRESS",
				":8080",
			),

			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},

		Database: DatabaseConfig{
			URL: getEnv(
				"DATABASE_URL",
				"postgres://postgres:postgres@localhost:5432/products?sslmode=disable",
			),

			MaxOpenConns: maxOpenConns,
			MaxIdleConns: maxIdleConns,

			ConnMaxLifetime: connMaxLifetime,
			ConnMaxIdleTime: connMaxIdleTime,
		},
	}, nil
}
