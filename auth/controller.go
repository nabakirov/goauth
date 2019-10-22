package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Index GET /
func Index(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.Header.Get("X-Original-URI"), "/")
	var token string
	if len(uri) < 2 {
		token = ""
	} else {
		token = uri[1]
	}
	var resp Response
	if val, ok := Tokens[token]; ok {
		if val == true {
			resp = Response{"ok", 200}
		} else {
			resp = Response{"invalid token", 403}
		}
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
