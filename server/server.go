package server

import (
	"fmt"
	baseconfig "github.com/blocktop/mp-common/config"
	"github.com/go-chi/chi"
	"github.com/stellar/go/support/config"
	"github.com/stellar/go/support/http"
	"log"
)

// RunHTTPServer runs an HTTP server with stellar-normative options configured.
func RunHTTPServer(h *chi.Mux) {
	cfg := baseconfig.GetConfig()

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
