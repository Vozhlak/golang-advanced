package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		logrus.WithFields(logrus.Fields{
			"status_code": wrapper.StatusCode,
			"method":      r.Method,
			"path":        r.URL.Path,
			"duration":    time.Since(start),
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		}).Info("HTTP request completed")
	})
}
