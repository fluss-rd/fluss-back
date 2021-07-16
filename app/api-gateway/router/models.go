package router

import (
	"context"
	"errors"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
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
type Router interface {
	Route() http.HandlerFunc
}

type Gateway struct {
	Router     Router
	Authorizer Authorizer
	Endpoint   Endpoint
}

func newRouter(endpoint Endpoint, mode TransportMode, rabbitClient rabbit.RabbitClient) (Router, error) {
	switch mode {
	case TransportModeHTTP:
		return newHttpRouter(endpoint), nil
	case TransportModeAMQP:
		return newRabbitMqRouter(endpoint, rabbitClient), nil
	}

	return nil, ErrInvalidTransportMode
}

func (g Gateway) handleEndpoint(ctx context.Context, endpoint Endpoint, requestHandler *mux.Router) {
	methods := []string{endpoint.Method}
	if endpoint.Method != http.MethodOptions {
		methods = append(methods, http.MethodOptions)
	}

	requestHandler.Handle(endpoint.Path, g.authMiddleware(ctx, g.Router.Route())).Methods(methods...)
}

func (g Gateway) authMiddleware(ctx context.Context, handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		setupPreflightResponse(&rw, r)
		if r.Method == http.MethodOptions {
			return
		}

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
