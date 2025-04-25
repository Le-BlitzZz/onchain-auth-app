package config

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/event"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
)

// log points to the global logger.
var log = event.Log

var c *Config

type Config struct {
	options *Options
	db      *gorm.DB
}

// NewConfig initialises a new configuration file.
func NewConfig(ctx *cli.Context) *Config {
	c = &Config{
		options: NewOptions(ctx),
	}
	return c
}

func (c *Config) Init() error {
	if err := c.connectDb(); err != nil {
		return err
	}

	// Show log message.
	log.Debug("Config: successfully initialized")

	return nil
}

// GetConfig returns app config.
func GetConfig() *Config {
	return c
}
