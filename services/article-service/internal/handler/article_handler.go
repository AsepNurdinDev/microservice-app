package handler

import (
	"net/http"
	"strconv"

	"article-service/internal/model"
	"article-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	svc *service.ArticleService
}

func NewArticleHandler(s *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: s}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var a model.Article
	c.BindJSON(&a)

	err := h.svc.Create(a)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

func (h *ArticleHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	skip := (page - 1) * limit

	data, _ := h.svc.GetAll(int64(skip), int64(limit))

	c.JSON(http.StatusOK, data)
}