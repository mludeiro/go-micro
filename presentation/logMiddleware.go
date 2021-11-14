package presentation

import (
	"go-micro/tools"
	"net/http"
	"time"
)

func LogMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		lrw := &logResponseWriter{ResponseWriter: w}

		next.ServeHTTP(lrw, r)

		tools.GetLogger().Printf(
			"%6dms %6s %6d %s",
			time.Since(startTime).Milliseconds(),
			r.Method,
			lrw.statusCode,
			r.URL.String())
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *logResponseWriter) Write(body []byte) (int, error) {
	return w.ResponseWriter.Write(body)
}
