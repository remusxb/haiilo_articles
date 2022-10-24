package model

import (
	"github.com/google/uuid"
	"github.com/remusxb/haiilo_articles/pkg/dto"
)

type Article struct {
	Id    uuid.UUID
	Title string
	Link  string
}

func (article Article) ToDTO() dto.ArticleOutput {
	return dto.ArticleOutput{
		Id:    article.Id,
		Title: article.Title,
		Link:  article.Link,
	}
}

func GetColumns() []string {
	return []string{
		"id",
		"title",
		"link",
	}
}

func GetTableName() string {
	return "articles"
}
