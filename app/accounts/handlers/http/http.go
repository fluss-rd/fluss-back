package httphandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/models"
	"github.com/flussrd/fluss-back/app/accounts/service"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/gorilla/mux"
)

var (
	// ErrMissingContentType missing content type
	ErrMissingContentType = httputils.NewBadRequestError("missing content type")
	// ErrInvalidBody invalid request body
	ErrInvalidBody = httputils.NewBadRequestError("invalid request body")
	// ErrMissingID missing id
	ErrMissingID = errors.New("missing id")
)

// HTTPHandler defines the functiosn that will handle http requests
type HTTPHandler interface {
	HandleCreateUser(ctx context.Context) http.HandlerFunc
	HandleGetUsers(ctx context.Context) http.HandlerFunc
	HandleUpdateUser(ctx context.Context) http.HandlerFunc
	HandleCreateRole(ctx context.Context) http.HandlerFunc
	HandleGetRoles(ctx context.Context) http.HandlerFunc
	HandleLogin(ctx context.Context) http.HandlerFunc
	HandleGetUser(ctx context.Context) http.HandlerFunc
}

type httpHandler struct {
	service service.Service
}

// NewHTTPHandler returns a new HTTPHandler object
func NewHTTPHandler(service service.Service) HTTPHandler {
	return httpHandler{
		service: service,
	}
}

// HandleCreateUser handles the create user request
func (h httpHandler) HandleCreateUser(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		user := models.User{}

		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		user, err = h.service.CreateUser(ctx, user)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusCreated, user)
	}
}

func (h httpHandler) HandleUpdateUser(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		userID := vars["id"]
		if userID == "" {
			httputils.RespondWithError(rw, ErrMissingID)
			return
		}

		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		patchRequest := httputils.PatchRequest{}
		err = json.NewDecoder(r.Body).Decode(&patchRequest)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		user, err := h.service.UpdateUser(ctx, patchRequest, userID)
		if err != nil {
			fmt.Println("updating_user_failed: " + err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, user)
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
			fmt.Println("creating_role_failed: " + err.Error())
			httputils.RespondWithError(rw, err)
			return
		}
	}
}

// HandleGetRoles handles the get roles request
func (h httpHandler) HandleGetRoles(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		roles, err := h.service.GetRoles(ctx)
		if err != nil {
			fmt.Println("getting_roles_failed: " + err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, roles)
	}
}

// HandleLogin handles the login http request
func (h httpHandler) HandleLogin(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		loginRequest := LoginRequestBody{}
		err = json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		response, err := h.service.Login(ctx, loginRequest.Email, loginRequest.Password)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, response)
	}
}

func (h httpHandler) HandleGetUser(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]

		if !ok {
			httputils.RespondWithError(rw, errors.New("missing id"))
			return
		}

		user, err := h.service.GetUser(ctx, id)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, user)
	}
}

func (h httpHandler) HandleGetUsers(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		users, err := h.service.GetUsers(ctx)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, users)
	}
}

func validateContentType(r http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return ErrMissingContentType
	}

	return nil
}
