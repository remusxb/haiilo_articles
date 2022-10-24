package dto

import (
	"net/url"

	"github.com/google/uuid"
)

type (
	CreateArticleInput struct {
		Title string `json:"title"`
		Link  string `json:"link"`
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

const (
	fieldTitle = "title"
	fieldLink  = "link"
)

func CreateArticlesInputFromUrlValues(queryParams url.Values) CreateArticleInput {
	return CreateArticleInput{
		Title: queryParams.Get(fieldTitle),
		Link:  queryParams.Get(fieldLink),
	}
}
