package main

import (
	"server/pkg/env"
)

type Config struct {
	portNumber string
}

func (c *Config) PortNumber() string {
	return c.portNumber
}

func NewConfig() *Config {
	return &Config{
		portNumber: env.MustGetEnvString("PORT", "8080"),
	}
}
