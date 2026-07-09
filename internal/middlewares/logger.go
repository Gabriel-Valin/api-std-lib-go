package middlewares

import (
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetLogger(r.Context()).Info("request started")
		next.ServeHTTP(w, r)
	})
}
