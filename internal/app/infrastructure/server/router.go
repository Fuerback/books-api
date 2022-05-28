package server

import "net/http"

type Router interface {
	GET(uri string, f func(resp http.ResponseWriter, r *http.Request))
	POST(uri string, f func(resp http.ResponseWriter, r *http.Request))
	PATCH(uri string, f func(resp http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(resp http.ResponseWriter, r *http.Request))
	Serve(port string)
}
