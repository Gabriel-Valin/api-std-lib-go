package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func Timing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		slog.Info(
			"request completed",
			"request_id", GetRequestID(r.Context()),
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"size_bytes", rw.size,
			"duration", time.Since(start),
		)
	})
}
