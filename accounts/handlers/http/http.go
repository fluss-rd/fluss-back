package httphandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/flussrd/fluss-back/accounts/models"
	"github.com/flussrd/fluss-back/accounts/service"
	"github.com/flussrd/fluss-back/accounts/shared/httputils"
)

var (
	// ErrMissingContentType missing content type
	ErrMissingContentType = httputils.NewBadRequestError("missing content type")
	// ErrInvalidBody invalid request body
	ErrInvalidBody = httputils.NewBadRequestError("invalid request body")
)

// HTTPHandler defines the functiosn that will handle http requests
type HTTPHandler interface {
	HandleCreateUser(ctx context.Context) http.HandlerFunc
	HandleCreateRole(ctx context.Context) http.HandlerFunc
}

type httpHandler struct {
	service service.Service
}

func NewHTTPHandler(service service.Service) HTTPHandler {
	return httpHandler{
		service: service,
	}
}

// HandleCreateUser handles the create user request
func (h httpHandler) HandleCreateUser(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

// HandleCreateRole handles the create role request
func (h httpHandler) HandleCreateRole(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		role := models.Role{}

		err = json.NewDecoder(r.Body).Decode(&role)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		err = h.service.CreateRole(ctx, role)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}
	}
}

func validateContentType(r http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		return ErrMissingContentType
	}

	return nil
}
