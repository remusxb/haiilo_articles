//go:build unit
// +build unit

package articles

import (
	"context"
	"testing"

	"github.com/remusxb/haiilo_articles/foundation/database/postgres"
	"github.com/remusxb/haiilo_articles/foundation/validator"
	"github.com/remusxb/haiilo_articles/internal/model"
	"github.com/remusxb/haiilo_articles/pkg/dto"
)

const (
	success    = "\033[32m✓"
	failed     = "\033[31m✗"
	colorReset = "\033[0m"
)

type repoMock struct {
	storage []dto.ArticleOutput
}

func (mock *repoMock) create(ctx context.Context, article model.Article) error {
	for _, a := range mock.storage {
		switch {
		case a.Title == article.Title:
			err := &postgres.ErrUniqueFieldViolation{
				Field: "title",
			}
			return err
		case a.Link == article.Link:
			err := &postgres.ErrUniqueFieldViolation{
				Field: "link",
			}
			return err
		}

	}

	mock.storage = append(mock.storage, dto.ArticleOutput{
		Id:    article.Id,
		Title: article.Title,
		Link:  article.Link,
	})

	return nil
}

func (mock *repoMock) list(ctx context.Context) (dto.ListArticlesOutput, error) {
	return dto.ListArticlesOutput{Articles: mock.storage}, nil
}

func TestHandler_Create(t *testing.T) {
	ctx := context.Background()
	validatorInstance := validator.GetInstance()
	storage := repoMock{
		[]dto.ArticleOutput{},
	}

	tcs := []struct {
		name    string
		service service
		input   dto.CreateArticleInput
		wantErr bool
		err     error
	}{
		{
			name: "successfully create article",
			service: service{
				validator:  validatorInstance,
				repository: &storage,
			},
			input: dto.CreateArticleInput{
				Title: "Test article",
				Link:  "/test-article",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "duplicate title",
			service: service{
				validator:  validatorInstance,
				repository: &storage,
			},
			input: dto.CreateArticleInput{
				Title: "Test article",
				Link:  "/test-article2",
			},
			wantErr: true,
			err:     &postgres.ErrUniqueFieldViolation{Field: "title"},
		},
		{
			name: "duplicate link",
			service: service{
				validator:  validatorInstance,
				repository: &storage,
			},
			input: dto.CreateArticleInput{
				Title: "Test article 2",
				Link:  "/test-article",
			},
			wantErr: true,
			err:     &postgres.ErrUniqueFieldViolation{Field: "link"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.service.create(ctx, tc.input)
			switch {
			case tc.wantErr && err.Error() != tc.err.Error():
				t.Errorf("%s Test failed. Want: %s Got: %s.\n%s", failed, tc.err, err, colorReset)
				return
			case !tc.wantErr && err != nil:
				t.Errorf("%sTest failed. No error expected; Got: %s.\n%s", failed, err, colorReset)
				return
			}

			t.Logf("%s Test `%s` passed.\n%s", success, tc.name, colorReset)
		})
	}
}
