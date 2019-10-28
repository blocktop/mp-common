package config

import (
	"github.com/caarlos0/env"
	"github.com/stellar/go/network"
)

// Config holds system configuration values.
type Config struct {
	Environment              string `env:"MP_ENV" envDefault:"local"`
	HorizonURL               string `env:"MP_HORIZON_URL" envDefault:"http://localhost:8000"`
	AnchorFile               string `env:"MP_ANCHOR_FILE"`
	WebAuthEndpoint          string `env:"MP_WEB_AUTH_ENDPOINT"`
	WebAuthPort              int    `env:"MP_WEB_AUTH_PORT" envDefault:"3000"`
	TransferServer           string `env:"MP_TRANSFER_SERVER"`
	TransferServerPort       int    `env:"MP_TRANSFER_SERVER_PORT" envDefault:"3001"`
	HTTPServerWriteTimeout   int    `env:"MP_HTTP_SERVER_WRITE_TIMEOUT" envDefault:"15"`
	HTTPServerReadTimeout    int    `env:"MP_HTTP_SERVER_READ_TIMEOUT" envDefault:"15"`
	HTTPServerIdleTimeout    int    `env:"MP_HTTP_SERVER_IDLE_TIMEOUT" envDefault:"60"`
	HTTPServerRequestTimeout int    `env:"MP_HTTP_SERVER_REQUEST_TIMEOUT" envDefault:"60"`
	HTTPServerMaxErrorPct    int    `env:"MP_HTTP_SERVER_MAX_ERROR_PCT" envDefault:"5"`
	HTTPServerHealthPath     string `env:"MP_HTTP_SERVER_HEALTH_PATH" envDefault:"/health"`
	TLSCertPath              string `env:"MP_TLS_CERT_PATH"`
	TLSKeyPath               string `env:"MP_TLS_KEY_PATH"`
	SigningKey               string `env:"MP_SIGNING_KEY"`
	SigningKeySeed           string `env:"MP_SIGNING_KEY_SEED"`
	AnchorName               string `env:"MP_ANCHOR_NAME" envDefault:"Blocktop"`
	NetworkPassphrase        string
}

const (
	Prod  = "prod"
	QA    = "qa"
	Dev   = "dev"
	Local = "local"
)

var cfg *Config

func init() {
	makeConfig()
}

func makeConfig() {
	cfg = &Config{}
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}

	if cfg.IsProduction() {
		cfg.NetworkPassphrase = network.PublicNetworkPassphrase
	} else {
		cfg.NetworkPassphrase = network.TestNetworkPassphrase
	}
}

// GetConfig creates a new Config and populatees it from environment variables.
func GetConfig() *Config {
	return cfg
}

func (c *Config) IsProduction() bool {
	return c.Environment == Prod
}

func (c *Config) IsQA() bool {
	return c.Environment == QA
}

func (c *Config) IsDev() bool {
	return c.Environment == Dev
}

func (c *Config) IsLocal() bool {
	return c.Environment == Local
}
