package middlewares

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

func newResponseWriter(
	w http.ResponseWriter,
) *responseWriter {

	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}
