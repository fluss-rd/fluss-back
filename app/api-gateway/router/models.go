package router

import (
	"context"
	"errors"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/gorilla/mux"
)

// TransportMode defines the way in which a request will be routed/proxied
type TransportMode string

const (
	// TransportModeHTTP transports mode http
	TransportModeHTTP TransportMode = "http"
	// TransportModeAMQP amqp
	TransportModeAMQP TransportMode = "amqp"
)

var (
	// TODO: consider: should we move all errors to a single errors file?
	ErrInvalidTransportMode = errors.New("invalid transport mode")
)

// routes requests
type RouterP interface {
	Route() http.HandlerFunc
}

type Gateway struct {
	Router     RouterP
	Authorizer Authorizer
	Endpoint   Endpoint
}

func newRouter(mode TransportMode) (RouterP, error) {
	switch mode {
	case TransportModeHTTP:
		return newHttpRouter(), nil
	}

	return nil, ErrInvalidTransportMode
}

func (g Gateway) handleEndpoint(ctx context.Context, endpoint Endpoint, requestHandler *mux.Router) {
	requestHandler.Handle(endpoint.Path, g.authMiddleware(ctx, g.Router.Route())).Methods(endpoint.Method)
}

func (g Gateway) authMiddleware(ctx context.Context, handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !g.Endpoint.Authorized {
			handler.ServeHTTP(rw, r)
			return
		}

		resource, err := getResourceFromEndpoint(g.Endpoint)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		action, err := getActionFromMethod(r.Method)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		sub, err := g.Authorizer.Authorize(ctx, r, resource, action)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		r.Header.Add("sub", sub)

		handler.ServeHTTP(rw, r)
	}
}
