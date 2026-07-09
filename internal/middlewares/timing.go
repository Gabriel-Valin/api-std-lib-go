package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Timing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		log.Printf(
			"request_id=%s method=%s path=%s status=%d size=%dB duration=%s",
			GetRequestID(r.Context()),
			r.Method,
			r.URL.Path,
			rw.statusCode,
			rw.size,
			time.Since(start),
		)
	})
}
