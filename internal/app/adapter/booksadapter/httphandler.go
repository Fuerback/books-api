package booksadapter

import "net/http"

type httpHandler struct{}

func NewHttpHandler() *httpHandler {
	return &httpHandler{}
}

func (c *httpHandler) Create(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Read(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Update(resp http.ResponseWriter, r *http.Request) {

}

func (c *httpHandler) Delete(resp http.ResponseWriter, r *http.Request) {

}
