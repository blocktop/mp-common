package server

import (
	"github.com/stellar/go/support/log"
	"net/http"
)

const (
	// BADTXN is the error code for a bad transaction from the client.
	BADTXN = "BADTXN"

	// SRVFAIL is the error code for a server side panic.
	SRVFAIL = "SRVFAIL"

	// BADACCT is the error code for a bad account string.
	BADACCT = "BADACCT"

	// CONTYP is the error code for bad content type.
	CONTYP = "CONTYP"

	// JWTSIG is the error code for JWT signing issues.
	JWTSIG = "JWTSIG"

	// AUTHHDR is the error code for incorrect Authorization header.
	AUTHHDR = "AUTHHDR"

	// JWTREQ is the error code for JWT token required.
	JWTREQ = "JWTREQ"

	// JWTINVD is the error code for an invalid JWT token.
	JWTINVD = "JWTINVD"

	// NOFILE is the error code for a file not found
	NOFILE = "NOFILE"
)

func ResponseError(w http.ResponseWriter, statusCode int, errorCode string, e error) {
	var msg string
	if statusCode >= 500 {
		msg = http.StatusText(statusCode)
		log.Errorf("internal error: %+v \n", e)
	} else {
		msg = e.Error()
	}

	data := map[string]interface{}{
		"status": statusCode,
	}
	if len(errorCode) > 0 {
		data["errorCode"] = errorCode
	}

	data["message"] = msg

	w.WriteHeader(statusCode)
	ResponseJSONMap(w, data)
}
