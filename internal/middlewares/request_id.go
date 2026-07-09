package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"sync/atomic"
)

type requestIDKey struct{}
type loggerKey struct{}

var (
	requestIDContextKey requestIDKey
	loggerContextKey    loggerKey
	requestIDCounter    uint64
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strconv.FormatUint(
			atomic.AddUint64(&requestIDCounter, 1),
			10,
		)

		requestLogger := slog.Default().With(
			"request_id", id,
			"method", r.Method,
			"path", r.URL.Path,
		)

		ctx := context.WithValue(
			r.Context(),
			requestIDContextKey,
			id,
		)

		ctx = context.WithValue(
			ctx,
			loggerContextKey,
			requestLogger,
		)

		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}

func GetRequestID(ctx context.Context) string {
	id, ok := ctx.Value(requestIDContextKey).(string)
	if !ok {
		return ""
	}

	return id
}

func GetLogger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerContextKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return logger
}
