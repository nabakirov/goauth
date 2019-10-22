package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Index GET /
func Index(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("X-Original-URI"), "/")[1]
	var resp Response
	if token == "asdf" {
		resp = Response{"ok", 200}
	} else {
		resp = Response{"invalid token", 403}
	}
	data, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.StatusCode)
	w.Write(data)
	return
}
