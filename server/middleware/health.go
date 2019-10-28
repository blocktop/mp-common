package middleware

import (
	"github.com/blocktop/mp-common/config"
	"github.com/blocktop/mp-common/server"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type healthWrapper struct {
	http.ResponseWriter
	statusCode int
}

var (
	startTime    int64
	requestCount int
	errorCount   int
)

func init() {
	startTime = time.Now().UnixNano()
}

// HealthHandler provides a API handler to retrieve health info from the server.
var HealthHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	var ok bool
	if requestCount < 50 {
		ok = true
	} else {
		errorPct := (errorCount/requestCount)*100
		ok = errorPct < cfg.HTTPServerMaxErrorPct
	}
	health := map[string]interface{}{
		"ok":             ok,
		"uptime_seconds": (time.Now().UnixNano() - startTime) / int64(time.Second),
		"numRequests":    requestCount,
		"numErrors":      errorCount,
	}

	server.ResponseJSONMap(w, health)
}

// HealthMiddleware is the outer most middle ware for the HTTP server.
var HealthMiddleware mux.MiddlewareFunc = func(next http.Handler) http.Handler {
	cfg := config.GetConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != cfg.HTTPServerHealthPath {
			requestCount++
		}

		wrapper := &healthWrapper{
			ResponseWriter: w,
			statusCode:     200,
		}

		next.ServeHTTP(wrapper, r)
	})
}

func (w *healthWrapper) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)

	if code >= 500 {
		errorCount++
	}
}
