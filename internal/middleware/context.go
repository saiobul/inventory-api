package middleware

import (
	"context"
	"net/http"
	"time"

	"inventory-api/internal/utils"

	"github.com/google/uuid"
)

func ContextMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}
			userID := r.Header.Get("X-User-ID")

			ctx = context.WithValue(ctx, utils.RequestIDKey, requestID)
			ctx = context.WithValue(ctx, utils.UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
