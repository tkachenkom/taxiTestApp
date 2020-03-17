package middlewares

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Logger middleware
func Logger(entry *logrus.Entry, durationThreshold time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			//TODO: ADD ID FOR EACH REQUEST
			defer func() {
				dur := time.Since(t1)
				lEntry := entry.WithFields(logrus.Fields{
					"path":     r.URL.Path,
					"duration": dur,
					"status":   ww.Status(),
				})
				lEntry.Debug("request finished")

				if dur > durationThreshold {
					lEntry.WithField("http_request", r).Warn("slow request")
				}
			}()

			entry.WithField("path", r.URL.Path).Debug("request started")
			next.ServeHTTP(ww, r)
		})
	}
}
