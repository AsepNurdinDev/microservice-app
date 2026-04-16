package service

import (
	"article-service/internal/model"
	"article-service/internal/repository"
)

type ArticleService struct {
	repo *repository.ArticleRepo
}

func NewArticleService(r *repository.ArticleRepo) *ArticleService {
	return &ArticleService{repo: r}
}

func (s *ArticleService) Create(a model.Article) error {
	return s.repo.Create(a)
}

func (s *ArticleService) GetAll(skip, limit int64) ([]model.Article, error) {
	return s.repo.FindAll(skip, limit)
}