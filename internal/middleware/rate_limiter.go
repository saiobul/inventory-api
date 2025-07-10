package middleware

import (
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiterMiddleware() func(http.Handler) http.Handler {
	// Define rate: 5 requests per minute
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  5,
	}

	// Use in-memory store
	store := memory.NewStore()

	// Create limiter instance
	instance := limiter.New(store, rate)

	// Create middleware
	middleware := stdlib.NewMiddleware(instance)

	return middleware.Handler
}
