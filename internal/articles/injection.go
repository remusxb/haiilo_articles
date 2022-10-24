package articles

import (
	"github.com/remusxb/haiilo_articles/foundation/database/postgres"
	"github.com/remusxb/haiilo_articles/foundation/validator"
)

func NewHandler() (Handler, error) {
	appValidator := validator.GetInstance()
	database, err := postgres.GetInstance()
	if err != nil {
		return Handler{}, err
	}

	repo := newRepository(database)
	srv := newService(appValidator, repo)
	handler := newHandler(srv)

	return handler, nil
}
