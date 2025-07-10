package middleware

import (
	"net/http"

	"inventory-api/internal/utils"

	"go.uber.org/zap"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ctx := r.Context()
				requestID := utils.GetRequestID(ctx)

				zap.L().Error("Recovered from panic",
					zap.Any("error", err),
					zap.String("request_id", requestID),
				)

				utils.RespondWithError(ctx, w, http.StatusInternalServerError, "Internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
