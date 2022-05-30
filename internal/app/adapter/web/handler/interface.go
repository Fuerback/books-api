package handler

import "net/http"

type BooksHandler interface {
	Create(resp http.ResponseWriter, r *http.Request)
	Read(resp http.ResponseWriter, r *http.Request)
	ReadBooks(resp http.ResponseWriter, r *http.Request)
	Update(resp http.ResponseWriter, r *http.Request)
	Delete(resp http.ResponseWriter, r *http.Request)
}
