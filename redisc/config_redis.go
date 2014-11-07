package redisc

import (
	"fmt"
)

type RedisConfig struct {
	Hostname string
	Port     string
}

func (c *RedisConfig) Connect_string() string {
	connect := fmt.Sprint(c.Hostname, ":", c.Port)
	return connect
}

func NewRedisConfig() *RedisConfig {
	cfg := &RedisConfig{
		Hostname: "localhost",
		Port:     "6379",
	}
	return cfg
}
