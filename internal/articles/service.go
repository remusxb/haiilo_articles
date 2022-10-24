package articles

import (
	"context"

	"github.com/google/uuid"
	"github.com/remusxb/haiilo_articles/foundation/validator"
	"github.com/remusxb/haiilo_articles/internal/model"
	"github.com/remusxb/haiilo_articles/pkg/dto"
)

type (
	service struct {
		validator  validator.AppValidatorInterface
		repository repositoryInterface
	}

	repositoryInterface interface {
		create(ctx context.Context, article model.Article) error
		list(ctx context.Context) (dto.ListArticlesOutput, error)
	}
)

func newService(
	validator validator.AppValidatorInterface,
	repository repositoryInterface,
) service {
	return service{
		validator:  validator,
		repository: repository,
	}
}

func (service service) create(ctx context.Context, input dto.CreateArticleInput) (dto.ArticleOutput, error) {
	if err := service.validator.Validate(input); err != nil {
		return dto.ArticleOutput{}, err
	}

	article := model.Article{
		Id:    uuid.New(),
		Title: input.Title,
		Link:  input.Link,
	}

	if err := service.repository.create(ctx, article); err != nil {
		return dto.ArticleOutput{}, err
	}

	return article.ToDTO(), nil
}

func (service service) list(ctx context.Context) (dto.ListArticlesOutput, error) {
	return service.repository.list(ctx)
}
