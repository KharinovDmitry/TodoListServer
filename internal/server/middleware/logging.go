package middleware

import (
	"fmt"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

func Logging(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, req)
			log.Info(fmt.Sprintf("method:%s URL:%s TIME:%s", req.Method, req.RequestURI, time.Since(start)))
		})
	}
}
