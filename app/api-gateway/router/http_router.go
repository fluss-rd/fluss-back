package router

import "net/http"

type httpRouter struct {
}

func newHttpRouter() RouterP {
	return httpRouter{}
}

func (router httpRouter) Route() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
