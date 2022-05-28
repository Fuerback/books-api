package server

import (
	"fmt"
	"github.com/Fuerback/books-api/internal/app/adapter/booksadapter"
	"github.com/Fuerback/books-api/internal/app/infrastructure/env"
	"net/http"
)

// HttpServer struct
type HttpServer struct {
	router  Router
	handler booksadapter.BooksHandler
}

// NewHttpServer New Server constructor
func NewHttpServer(handler booksadapter.BooksHandler) *HttpServer {
	return &HttpServer{router: NewMuxRouter(), handler: handler}
}

func (s *HttpServer) Run() {
	s.router.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})

	s.router.Serve(env.GetEnvWithDefaultAsString("PORT", ":8080"))
}
