package config

import (
	"github.com/caarlos0/env"
	"github.com/stellar/go/network"
)

type Config interface {
	Parse()
	IsDev() bool
	IsQA() bool
	IsProduction() bool
	IsLocal() bool
}

// BaseConfig holds system configuration values.
type BaseConfig struct {
	Environment              string `env:"MP_ENV" envDefault:"local"`
	HTTPServerWriteTimeout   int    `env:"MP_HTTP_SERVER_WRITE_TIMEOUT" envDefault:"15"`
	HTTPServerReadTimeout    int    `env:"MP_HTTP_SERVER_READ_TIMEOUT" envDefault:"15"`
	HTTPServerIdleTimeout    int    `env:"MP_HTTP_SERVER_IDLE_TIMEOUT" envDefault:"60"`
	HTTPServerRequestTimeout int    `env:"MP_HTTP_SERVER_REQUEST_TIMEOUT" envDefault:"60"`
	HTTPServerMaxErrorPct    int    `env:"MP_HTTP_SERVER_MAX_ERROR_PCT" envDefault:"5"`
	HTTPServerHealthPath     string `env:"MP_HTTP_SERVER_HEALTH_PATH" envDefault:"/health"`
	NetworkPassphrase        string
}

const (
	Prod  = "prod"
	QA    = "qa"
	Dev   = "dev"
	Local = "local"
)

func (c *BaseConfig) Parse() {
	err := env.Parse(c)
	if err != nil {
		panic(err)
	}

	if c.IsProduction() {
		c.NetworkPassphrase = network.PublicNetworkPassphrase
	} else {
		c.NetworkPassphrase = network.TestNetworkPassphrase
	}
}

func (c *BaseConfig) IsProduction() bool {
	return c.Environment == Prod
}

func (c *BaseConfig) IsQA() bool {
	return c.Environment == QA
}

func (c *BaseConfig) IsDev() bool {
	return c.Environment == Dev
}

func (c *BaseConfig) IsLocal() bool {
	return c.Environment == Local
}
