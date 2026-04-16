package repository

import (
	"context"
	"article-service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleRepo struct {
	col *mongo.Collection
}

func NewArticleRepo(db *mongo.Database) *ArticleRepo {
	return &ArticleRepo{col: db.Collection("articles")}
}

func (r *ArticleRepo) Create(a model.Article) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	_, err := r.col.InsertOne(context.Background(), a)
	return err
}

func (r *ArticleRepo) FindAll(skip, limit int64) ([]model.Article, error) {
    ctx := context.Background()

    cursor, err := r.col.Find(ctx, bson.M{}, &options.FindOptions{
        Skip:  &skip,
        Limit: &limit,
    })
    if err != nil {
        return nil, err
    }

    var result []model.Article
    err = cursor.All(ctx, &result)
    return result, err
}