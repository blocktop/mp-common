package server

import (
	"fmt"
	config2 "github.com/blocktop/mp-common/config"
	"github.com/go-chi/chi"
	"github.com/stellar/go/support/config"
	"github.com/stellar/go/support/http"
	"log"
)

func RunHTTPServer(h *chi.Mux, name string, port int) {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", port)

	cfg := config2.GetConfig()
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

	log.Printf("starting %s at: %s", name, listenAddr)

	http.Run(httpCfg)

	log.Printf("shutting down %s", name)
}
