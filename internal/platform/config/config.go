package config

import "os"

type Config struct {
	ServiceName string
	HTTPAddress string
	DatabaseURL string
	RedisAddr   string
	NATSURL     string
	LogLevel    string
}

func Load() Config {
	return Config{
		ServiceName: get("SERVICE_NAME", "go-microservice-template"),
		HTTPAddress: get("HTTP_ADDRESS", ":3002"),
		DatabaseURL: get("DATABASE_URL", ""),
		RedisAddr:   get("REDIS_ADDR", ""),
		NATSURL:     get("NATS_URL", ""),
		LogLevel:    get("LOG_LEVEL", "debug"),
	}
}

func get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
