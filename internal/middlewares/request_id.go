package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"sync/atomic"
)

type contextKey string

const requestIDKey contextKey = "request_id"

var requestIDCounter uint64

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strconv.FormatUint(
			atomic.AddUint64(&requestIDCounter, 1),
			10,
		)
		ctx := context.WithValue(
			r.Context(),
			requestIDKey,
			id,
		)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}
	return id
}
