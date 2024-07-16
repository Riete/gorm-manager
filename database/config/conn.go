package config

import (
	"time"
)

const (
	DefaultMaxConn         = 20
	DefaultMaxConnLifetime = 3 * time.Minute
)

type ConnConfig struct {
	MaxConn         int
	MaxConnLifetime time.Duration
}
