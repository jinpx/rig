package redis

import (
	"log"
	"time"
)

// Config ...
type Config struct {
	Name         string        `json:"name" yaml:"name"`
	Host         string        `json:"host" yaml:"host"`
	Username     string        `json:"username" yaml:"username"`
	Password     string        `json:"password" yaml:"password"`
	PoolSize     int           `json:"poolSize" yaml:"poolSize"`
	MaxRetries   int           `json:"maxRetries" yaml:"maxRetries"`
	MinIdleConns int           `json:"minIdleConns" yaml:"minIdleConns"`
	DialTimeout  time.Duration `json:"dialTimeout" yaml:"dialTimeout"`
	ReadTimeout  time.Duration `json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout" yaml:"writeTimeout"`
	IdleTimeout  time.Duration `json:"idleTimeout" yaml:"idleTimeout"`
	logger       *log.Logger
}
