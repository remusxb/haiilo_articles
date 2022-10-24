package articles

import (
	"context"
	"log"
	"net/http"

	"github.com/remusxb/haiilo_articles/foundation/http/response"
	"github.com/remusxb/haiilo_articles/pkg/dto"
)

type (
	Handler struct {
		service service
	}
)

func newHandler(service service) Handler {
	return Handler{service: service}
}

func (handler Handler) Create(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	var output interface{}
	statusCode := http.StatusCreated
	queryParams := request.URL.Query()
	input := dto.CreateArticlesInputFromUrlValues(queryParams)

	output, err := handler.service.create(ctx, input)
	if err != nil {
		// some error handling
		statusCode = http.StatusInternalServerError
		output = err
	}

	_, err = response.Write(writer, output, statusCode)
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

func (handler Handler) List(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	statusCode := http.StatusOK

	_, err := response.Write(writer, nil, statusCode)
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}
