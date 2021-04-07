package httphandler

import (
	"net/http"

	"github.com/flussrd/fluss-back/accounts/service"
)

// HTTPHandler defines the functiosn that will handle http requests
type HTTPHandler interface {
	HandleCreateUser() http.HandlerFunc
	HandleCreateRole() http.HandlerFunc
}

type httpHandler struct {
	service service.Service
}

// HandleCreateUser handles the create user request
func (h httpHandler) HandleCreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

// HandleCreateRole handles the create role request
func (h httpHandler) HandleCreateRole() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
