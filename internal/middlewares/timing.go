package middlewares

import (
	"net/http"
	"time"
)

func Timing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		GetLogger(r.Context()).Info(
			"request completed",
			"status", rw.statusCode,
			"size_bytes", rw.size,
			"duration", time.Since(start),
		)
	})
}
