package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	Port               string
	RedisAddr          string
	RedisPassword      string
	RedisDB            int
	CloudWatchLogGroup string
	CloudWatchStream   string
	ProfilingAddr      string
}

func LoadConfig() (*Config, error) {
	redisDB, err := getEnvAsInt("REDIS_DB", 0)
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB: %w", err)
	}

	return &Config{
		DBHost:             getEnv("DB_HOST", ""),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", ""),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", ""),
		Port:               getEnv("PORT", "8080"),
		RedisAddr:          getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:      getEnv("REDIS_PASSWORD", ""),
		RedisDB:            redisDB,
		CloudWatchLogGroup: getEnv("CLOUDWATCH_LOG_GROUP", "InventoryLogGroup"),
		CloudWatchStream:   getEnv("CLOUDWATCH_STREAM", "AppStream"),
		ProfilingAddr:      getEnv("PROFILING_ADDR", "localhost:6060"),
	}, nil
}

// Helper to read env with default fallback
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// Helper to read int env with default fallback
func getEnvAsInt(key string, defaultVal int) (int, error) {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal, nil
	}
	return strconv.Atoi(valStr)
}
