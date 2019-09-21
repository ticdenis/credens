package middleware

import (
	"credens/src/infrastructure/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewLoggingMiddleware(logger logging.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				logger.Log(fmt.Sprintf("%s %s", r.Method, r.RequestURI))

				next.ServeHTTP(w, r)
			},
		)
	}
}
