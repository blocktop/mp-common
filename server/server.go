package server

import (
	"fmt"
	baseconfig "github.com/blocktop/mp-common/config"
	"github.com/go-chi/chi"
	"github.com/stellar/go/support/config"
	"github.com/stellar/go/support/http"
	"log"
)

func RunHTTPServer(h *chi.Mux, c baseconfig.Config) {
	cfg := c.(*baseconfig.BaseConfig)

	listenAddr := fmt.Sprintf("0.0.0.0:%d", cfg.HTTPServerPort)

	tls := &config.TLS{
		CertificateFile: cfg.TLSCertPath,
		PrivateKeyFile:  cfg.TLSKeyPath,
	}

	httpCfg := http.Config{
		Handler:             h,
		ListenAddr:          listenAddr,
		TLS:                 tls,
		ShutdownGracePeriod: 15,
	}

	log.Printf("starting %s at: %s", cfg.HTTPServerName, listenAddr)

	http.Run(httpCfg)

	log.Printf("shutting down %s", cfg.HTTPServerName)
}
