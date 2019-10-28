package server

import (
	"encoding/json"
	"net/http"
)

const (
	DefaultJSONResponse = `{"result":"ok"}`
)

func ResponseJSONString(w http.ResponseWriter, j string) {
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write([]byte(j))
	if err != nil {
		_, _ = w.Write([]byte(`"failed to encode json data"`))
	}
}

func ResponseJSONMap(w http.ResponseWriter, j map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		_, _ = w.Write([]byte(`"failed to encode json data"`))
	}
}
