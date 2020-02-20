package config

import (
	"github.com/x-punch/go-config"
)

// Config represents the server configuration.
type Config struct {
	Address string `toml:"address"`
}

// NewConfig builds a new configuration with default values.
func NewConfig() *Config {
	return &Config{Address: ":80"}
}

// Load parse config info from config file and env args
func (c *Config) Load() error {
	return config.Load(c, "config.toml")
}
