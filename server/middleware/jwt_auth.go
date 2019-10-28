package middleware

import (
	"github.com/blocktop/mp-common/config"
	"github.com/blocktop/mp-common/server"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// JWTAuthMiddleWare is the outer most middle ware for the HTTP server.
var JWTAuthMiddleWare mux.MiddlewareFunc = func(next http.Handler) http.Handler {
	cfg := config.GetConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sToken string
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) > 0 {
			if !strings.HasPrefix(authHeader, "Bearer ") {
				responseAuthRequired(w)
				return
			}

			sToken = authHeader[7:]
		} else {
			sToken = r.URL.Query().Get("jwt")
		}

		if len(sToken) == 0 {
			responseAuthRequired(w)
			return
		}

		token, err := jwt.ParseWithClaims(sToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.SigningKeySeed), nil
		})
		if err != nil {
			responseAuthRequired(w)
			return
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			r.Header.Add(server.SubjectAccountIDHeader, claims.Subject)
			next.ServeHTTP(w, r)
		} else {
			responseAuthRequired(w)
		}
	})
}

func responseAuthRequired(w http.ResponseWriter) {
	const data = `{"type": "authentication_required"}`
	w.WriteHeader(http.StatusForbidden)
	server.ResponseJSONString(w, data)
}