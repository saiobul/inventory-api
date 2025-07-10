package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initializes the global logger with the specified log level.
// Supported levels: "debug", "info", "warn", "error"
func InitLogger(level string) {
	var cfg zap.Config

	switch level {
	case "debug":
		cfg = zap.NewDevelopmentConfig()
	case "info", "warn", "error":
		cfg = zap.NewProductionConfig()
		cfg.Level = parseLevel(level)
	default:
		cfg = zap.NewProductionConfig()
	}

	var err error
	Log, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

// parseLevel maps string level to zapcore.Level
func parseLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
