package model

import "time"

type Article struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Content   string    `bson:"content" json:"content"`
	AuthorID  string    `bson:"author_id" json:"author_id"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}