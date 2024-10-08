package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Token string `env:"TOKEN" envDefault:""`
}

// Builder defines the builder for the Config struct
type Builder struct {
	cfg *Config
}

// NewConfigBuilder initializes the ConfigBuilder with default values
func NewConfigBuilder() *Builder {
	return &Builder{
		cfg: &Config{},
	}
}

// FromEnv parses environment variables into the ConfigBuilder
func (b *Builder) FromEnv() *Builder {
	if err := env.Parse(b.cfg); err != nil {
		fmt.Printf("Error parsing environment variables: %+v\n", err)
	}
	return b
}

// Build returns the final configuration
func (b *Builder) Build() *Config {
	return b.cfg
}
