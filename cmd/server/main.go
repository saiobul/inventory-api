package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"inventory-api/api"
	"inventory-api/internal/cache"
	"inventory-api/internal/config"
	"inventory-api/internal/db"
	"inventory-api/internal/logger"
	"inventory-api/internal/middleware"
	"inventory-api/internal/product"
	"inventory-api/pkg/aws"
)

func main() {
	// Initialize logger
	logger.InitLogger("debug")
	logger.Log.Info("Logger initialized")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize optional CloudWatch logger
	if cwLogger, err := aws.NewCloudWatchLogger(cfg.CloudWatchLogGroup, cfg.CloudWatchStream); err != nil {
		logger.Log.Warn("CloudWatch logger init failed", zap.Error(err))
	} else {
		cwLogger.SendLog("CloudWatch logger initialized successfully")
	}

	// Initialize Redis
	redisClient := cache.NewRedisClient(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if _, err := redisClient.Get("healthcheck"); err != nil {
		logger.Log.Info("Redis connection check passed (expected miss)")
	} else {
		logger.Log.Info("Redis is up and running")
	}

	// Initialize database
	dbConn, err := db.ConnectPostgres(cfg)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Setup product service
	repo := product.NewRepository(dbConn)
	service := product.NewService(repo, redisClient)
	handler := product.NewHandler(service)

	// Setup routes and middleware
	router := api.SetupRoutes(handler)
	router = applyMiddlewares(router)

	// Start profiling server
	go func() {
		logger.Log.Info("Profiling available at http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe(cfg.ProfilingAddr, nil); err != nil {
			logger.Log.Fatal("Profiling server failed", zap.Error(err))
		}
	}()

	// Graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		logger.Log.Info("Server starting", zap.String("port", cfg.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server failed", zap.Error(err))
		}
	}()

	<-stop
	logger.Log.Info("Shutting down server...")
	// Add cleanup logic here if needed
}

func applyMiddlewares(h http.Handler) http.Handler {
	h = middleware.CORSMiddleware(h)
	h = middleware.VersionValidatorMiddleware([]string{"v1"})(h)
	h = middleware.RateLimiterMiddleware()(h)
	h = middleware.RecoveryMiddleware(h)
	h = middleware.ContextMiddleware(5 * time.Second)(h)
	h = middleware.LoggingMiddleware(h)
	return h
}
