package server

import (
	"fmt"
	"net/http"

	"github.com/Fuerback/books-api/internal/app/adapter/web/handler"
	"github.com/Fuerback/books-api/internal/app/infrastructure/env"
)

// HttpServer struct
type HttpServer struct {
	router  Router
	handler handler.BooksHandler
}

// NewHttpServer New Server constructor
func NewHttpServer(handler handler.BooksHandler) *HttpServer {
	return &HttpServer{router: newMuxRouter(), handler: handler}
}

func (s *HttpServer) Run() {
	s.router.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})

	s.router.POST("/v1/books", func(resp http.ResponseWriter, r *http.Request) {
		s.handler.Create(resp, r)
	})

	s.router.GET("/v1/books", func(resp http.ResponseWriter, r *http.Request) {
		s.handler.ReadBooks(resp, r)
	})

	s.router.PATCH("/v1/books/{id}", func(resp http.ResponseWriter, r *http.Request) {
		s.handler.Update(resp, r)
	})

	s.router.GET("/v1/books/{id}", func(resp http.ResponseWriter, r *http.Request) {
		s.handler.Read(resp, r)
	})

	s.router.DELETE("/v1/books/{id}", func(resp http.ResponseWriter, r *http.Request) {
		s.handler.Delete(resp, r)
	})

	s.router.Serve(env.GetEnvWithDefaultAsString("PORT", ":8080"))
}
