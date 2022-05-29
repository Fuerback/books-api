package handler

import (
	"github.com/Fuerback/books-api/internal/app/domain"
	"net/http"
)

type httpHandler struct {
	bookService domain.Book
}

func NewHttpHandler(bookService domain.Book) BooksHandler {
	return &httpHandler{bookService: bookService}
}

func (c *httpHandler) Create(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Read(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Update(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Delete(resp http.ResponseWriter, r *http.Request) {

}
