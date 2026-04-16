package main

import (
	"os"

	"article-service/internal/config"
	"article-service/internal/handler"
	"article-service/internal/middleware"
	"article-service/internal/repository"
	"article-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectMongo(os.Getenv("MONGO_URI"))

	repo := repository.NewArticleRepo(db)
	svc := service.NewArticleService(repo)
	h := handler.NewArticleHandler(svc)

	r := gin.Default()

	article := r.Group("/articles")
	article.Use(middleware.JWTAuth())
	{
		article.POST("/", h.Create)
		article.GET("/", h.GetAll)
	}

	r.Run(":8002")
}