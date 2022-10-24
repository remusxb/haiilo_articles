package routing

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type AppRouter struct {
	ctx         context.Context
	mux         *mux.Router
	middlewares []Middleware
}

type Middleware func(Handler) Handler
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func NewAppRouter(ctx context.Context, mux *mux.Router, middlewares []Middleware) *AppRouter {
	return &AppRouter{ctx, mux, middlewares}
}

func (router AppRouter) SubRouter(prefix string, middlewares ...Middleware) *AppRouter {
	newRouter := router.mux.PathPrefix(prefix).Subrouter()
	newMiddlewares := append(router.middlewares, middlewares...)
	return NewAppRouter(router.ctx, newRouter, newMiddlewares)
}

func (router AppRouter) Handle(method, path string, handler Handler, middlewares ...Middleware) {
	handler = router.wrap(middlewares, handler)
	handler = router.wrap(router.middlewares, handler)
	fn := func(writer http.ResponseWriter, request *http.Request) {
		err := handler(router.ctx, writer, request)
		if err != nil {
			return
		}
	}

	router.mux.HandleFunc(path, fn).Methods(method)
}

func (router AppRouter) wrap(mw []Middleware, handler Handler) Handler {
	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
