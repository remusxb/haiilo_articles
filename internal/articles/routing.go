package articles

import (
	"log"
	"net/http"

	"github.com/remusxb/haiilo_articles/foundation/routing"
)

func NewRouter(appRouter *routing.AppRouter) {
	handler, err := NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	subRouter := appRouter.SubRouter("/articles")
	subRouter.Handle(http.MethodGet, "", handler.List)
	subRouter.Handle(http.MethodGet, "", handler.Create)
}
