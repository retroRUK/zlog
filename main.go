package zlog

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func HttpError(w http.ResponseWriter, msg string, err error, code int) {
	slog.Error(msg, "error", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func Error(msg string, err error) {
	slog.Error(msg, "error", err)
}
