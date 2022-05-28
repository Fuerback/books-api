package booksadapter

import (
	"github.com/Fuerback/books-api/internal/app/domain/book"
	"net/http"
)

type httpHandler struct {
	bookService book.Book
}

func NewHttpHandler(bookService book.Book) BooksHandler {
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
