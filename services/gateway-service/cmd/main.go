package main

import (
	"log"
	"os"

	"gateway-service/internal/handler"
	"gateway-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	h := handler.NewGatewayHandler() 

	// PUBLIC
	r.Any("/auth/*path", h.AuthProxy())

	// PROTECTED
	r.Any("/articles/*path", middleware.JWTAuth(), h.ArticleProxy())

	port := getEnv("PORT", "8080")
	log.Printf("[GATEWAY] Running on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("[GATEWAY] Failed to start: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}