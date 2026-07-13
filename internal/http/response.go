package http

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	v any,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) error {
	return writeJSON(
		w,
		status,
		errorResponse{
			Error: message,
		},
	)
}
