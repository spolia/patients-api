package web

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest  = Error{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrInvalidJSON = Error{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
)

type Error struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e *Error) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
