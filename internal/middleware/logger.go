package middleware

import (
	"net/http"
	"time"

	"inventory-api/internal/logger"
	"inventory-api/pkg/aws"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func getRequestID(r *http.Request) string {
	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = uuid.New().String()
	}
	return reqID
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		path := r.URL.Path
		requestID := getRequestID(r)
		userID := r.Header.Get("X-User-ID")

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		// Send log to CloudWatch
		cwLogger, err := aws.NewCloudWatchLogger("InventoryLogGroup", "AppStream")
		if err == nil {
			cwLogger.SendLog("Started " + method + " " + path + " [RequestID: " + requestID + "]")
		}

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)
		status := recorder.status

		fields := []zap.Field{
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("duration", duration),
			zap.String("request_id", requestID),
		}
		if userID != "" {
			fields = append(fields, zap.String("user_id", userID))
		}

		switch {
		case status >= 500:
			logger.Log.Error("Request completed", fields...)
		case status >= 400:
			logger.Log.Warn("Request completed", fields...)
		default:
			logger.Log.Info("Request completed", fields...)
		}

		if err == nil {
			cwLogger.SendLog("Completed " + method + " " + path + " with status " + http.StatusText(status) + " in " + duration.String() + " [RequestID: " + requestID + "]")
		}
	})
}
