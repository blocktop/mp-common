package server

import "net/http"

const (
	SubjectAccountIDHeader = "X-Subject-AccountID"
)

func SubjectAccountID(r *http.Request) string {
	// this is set in middleware/jwt_auth.go
	return r.Header.Get(SubjectAccountIDHeader)
}
