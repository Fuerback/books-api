package booksadapter

import "net/http"

type BooksHandler interface {
	Create(resp http.ResponseWriter, r *http.Request)
	Get(resp http.ResponseWriter, r *http.Request)
	Update(resp http.ResponseWriter, r *http.Request)
	Delete(resp http.ResponseWriter, r *http.Request)
}
