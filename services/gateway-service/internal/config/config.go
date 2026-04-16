package config

import "os"

type Config struct {
	AuthService    string
	ArticleService string
	JWTSecret      string
}

func Load() Config {
	return Config{
		AuthService:    getEnv("AUTH_SERVICE", "http://auth-service:8001"),
		ArticleService: getEnv("ARTICLE_SERVICE", "http://article-service:8002"),
		JWTSecret:      getEnv("JWT_SECRET", "supersecret"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}