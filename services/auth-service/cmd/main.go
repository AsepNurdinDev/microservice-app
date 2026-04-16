package main

import (
	"log"
	nethttp "net/http"

	"auth-service/internal/config"
	httphandler "auth-service/internal/delivery/http"
	"auth-service/internal/infrastructure/database"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env (hanya untuk lokal, di Docker pakai environment di compose)
	if err := godotenv.Load(); err != nil {
		log.Println("[CONFIG] No .env file found, using environment variables")
	}

	// load & validasi config
	cfg := config.Load()

	// set gin mode ke release (matikan debug log yang verbose)
	gin.SetMode(gin.ReleaseMode)

	// koneksi database
	db, err := database.NewPostgres(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	if err != nil {
		log.Fatalf("[DB] Failed to connect: %v", err)
	}
	defer db.Close()
	log.Println("[DB] Connected successfully")

	// dependency injection
	userRepo := database.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)
	handler := httphandler.NewHandler(authUsecase)

	// router
	r := gin.New()
	r.Use(gin.Recovery()) // recover dari panic
	r.Use(gin.Logger())   // request logging

	// health check — dipakai Docker healthcheck & gateway depends_on
	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(nethttp.StatusServiceUnavailable, gin.H{
				"status":  "unhealthy",
				"service": "auth-service",
				"error":   "database unreachable",
			})
			return
		}
		c.JSON(nethttp.StatusOK, gin.H{
			"status":  "ok",
			"service": "auth-service",
		})
	})

	// routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	log.Printf("[AUTH] Running on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("[AUTH] Failed to start: %v", err)
	}
}