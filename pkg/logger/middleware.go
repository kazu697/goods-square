package logger

import (
	"net/http"

	"go.uber.org/zap"
)

type RequestLogger struct {
	logger *zap.Logger
}

func (rl *RequestLogger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
	})
}
