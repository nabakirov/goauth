package auth

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

var Tokens map[string]bool
