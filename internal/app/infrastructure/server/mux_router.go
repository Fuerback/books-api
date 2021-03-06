package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func newMuxRouter() Router {
	return &muxRouter{}
}

func (r *muxRouter) GET(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (r *muxRouter) POST(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (r *muxRouter) PATCH(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PATCH")
}

func (r *muxRouter) DELETE(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (r *muxRouter) Serve(port string) {
	fmt.Printf("Mux running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
