package dto

import (
	"github.com/google/uuid"
)

type (
	CreateArticleInput struct {
		Title string `json:"title" validate:"required,min=5,max=500"`
		Link  string `json:"link" validate:"required,min=2,max=1000"`
	}

	ArticleOutput struct {
		Id    uuid.UUID `json:"id"`
		Title string    `json:"title"`
		Link  string    `json:"link"`
	}

	ListArticlesOutput struct {
		Articles []ArticleOutput `json:"articles"`
	}
)
