package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type contextKey string

const (
	RequestIDKey contextKey = "requestID"
	UserIDKey    contextKey = "userID"
)

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(RequestIDKey).(string); ok {
		return v
	}
	return ""
}
func GetUserID(ctx context.Context) string {
	if v, ok := ctx.Value(UserIDKey).(string); ok {
		return v
	}
	return ""
}

func RespondWithError(ctx context.Context, w http.ResponseWriter, code int, message string) {
	requestID := GetRequestID(ctx)
	zap.L().Error("Responding with error",
		zap.String("request_id", requestID),
		zap.Int("status_code", code),
		zap.String("error", message),
	)

	RespondWithJSON(ctx, w, code, APIResponse{
		Success: false,
		Error:   message,
	})
}

func RespondWithJSON(ctx context.Context, w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		requestID := GetRequestID(ctx)
		zap.L().Error("Failed to encode JSON response",
			zap.String("request_id", requestID),
			zap.Error(err),
		)
	}
}
