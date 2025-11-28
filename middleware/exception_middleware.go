package mdw

import (
	"log/slog"
	"net/http"
)

func ExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Exeption:", slog.AnyValue(err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
