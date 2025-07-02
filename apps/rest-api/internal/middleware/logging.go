package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the response writer to capture status code
		wrappedWriter := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)

		log.Printf(
			"%s %s %d %s %s",
			r.Method,
			r.RequestURI,
			wrappedWriter.status,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
