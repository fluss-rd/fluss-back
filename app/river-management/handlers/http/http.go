package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/river-management/models"
	"github.com/flussrd/fluss-back/app/river-management/service"
	"github.com/gorilla/mux"
)

var (
	// ErrMissingContentType missing content type
	ErrMissingContentType = httputils.NewBadRequestError("missing content type")
	// ErrInvalidBody invalid request body
	ErrInvalidBody = httputils.NewBadRequestError("invalid request body")
	// ErrMissingSub missing sub
	ErrMissingSub = errors.New("missing sub")
)

// HTTPHandler defines the methods that will handle the incoming http requests
type HTTPHandler interface {
	HandleCreateRiver(ctx context.Context) http.HandlerFunc
	HandleGetRivers(ctx context.Context) http.HandlerFunc
	HandleGetRiver(ctx context.Context) http.HandlerFunc

	HandleCreateModule(ctx context.Context) http.HandlerFunc
	HandleGetModule(ctx context.Context) http.HandlerFunc
	HandleUpdateModule(ctx context.Context) http.HandlerFunc
	HandleGetModules(ctx context.Context) http.HandlerFunc
}

type httpHandler struct {
	s service.Service
}

// NewHTTPHandler returns a new httpHandler entity that will handle requests
func NewHTTPHandler(s service.Service) HTTPHandler {
	return httpHandler{
		s: s,
	}
}

// HandleCreateRiver handles the create river http request
func (h httpHandler) HandleCreateRiver(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		river := models.River{}
		err = json.NewDecoder(r.Body).Decode(&river)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		sub := r.Header.Get("sub")
		if sub == "" {
			httputils.RespondWithError(rw, ErrMissingSub)
			return
		}

		river.UserID = sub

		river, err = h.s.CreateRiver(ctx, river)
		if err != nil {
			fmt.Println("creating_river_failed: " + err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusCreated, river)
	}
}

// HandleGetRivers handles the get rivers http request
func (h httpHandler) HandleGetRivers(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rivers, err := h.s.GetRiversN(ctx)
		if err != nil {
			fmt.Println("fetching_rivers_failed: ", err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, rivers)
	}
}

// HandlGetRiver handles the get river http request
func (h httpHandler) HandleGetRiver(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		riverID := mux.Vars(r)["id"]

		river, err := h.s.GetRiver(ctx, riverID)
		if err != nil {
			fmt.Println("fetching_single_river_failed: ", err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, river)
	}
}

// HandleCreateModule handles the create module http request
func (h httpHandler) HandleCreateModule(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := validateContentType(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		module := models.Module{}
		err = json.NewDecoder(r.Body).Decode(&module)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		sub := r.Header.Get("sub")
		if sub == "" {
			httputils.RespondWithError(rw, ErrMissingSub)
			return
		}

		module.UserID = sub

		module, err = h.s.CreateModule(ctx, module)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusCreated, module)
	}
}

func (h httpHandler) HandleUpdateModule(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		options := models.ModuleUpdateOptions{}
		err := json.NewDecoder(r.Body).Decode(&options)
		if err != nil {
			httputils.RespondWithError(rw, ErrInvalidBody)
			return
		}

		_, err = h.s.UpdateModule(ctx, id, options)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusCreated, map[string]string{"message": "updated"})
	}
}

// HandleGetModule handle get module handles the get module http request
func (h httpHandler) HandleGetModule(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		module, err := h.s.GetModule(ctx, id)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, module)
	}
}

// HandleGetModules handles the GET modules http request
func (h httpHandler) HandleGetModules(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		modules, err := h.s.GetModulesN(ctx)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, modules)
	}
}

func validateContentType(r http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return ErrMissingContentType
	}

	return nil
}
