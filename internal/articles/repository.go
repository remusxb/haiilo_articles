package articles

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/remusxb/haiilo_articles/internal/model"
	"github.com/remusxb/haiilo_articles/pkg/dto"
)

type repository struct {
	database *sqlx.DB
}

func newRepository(database *sqlx.DB) repository {
	return repository{
		database: database,
	}
}

func (repository repository) create(ctx context.Context, article model.Article) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(%s) VALUES($1, $2, $3)`,
		model.GetTableName(),
		strings.Join(model.GetColumns(), ","),
	)

	_, err := repository.database.ExecContext(ctx, query, article.Id, article.Title, article.Link)
	if err != nil {
		return err
	}

	return nil
}

func (repository repository) list(ctx context.Context) (dto.ListArticlesOutput, error) {
	var output dto.ListArticlesOutput
	query := fmt.Sprintf(
		`SELECT %s FROM %s`,
		strings.Join(model.GetColumns(), ","),
		model.GetTableName(),
	)

	err := repository.database.SelectContext(ctx, &output, query)
	if err != nil {
		return dto.ListArticlesOutput{}, err
	}

	return output, err
}
