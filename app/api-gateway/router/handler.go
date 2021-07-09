package router

import (
	"context"

	"github.com/gorilla/mux"
)

type Handler interface {
	handleEndpoints(ctx context.Context) error
}

// TODO: consider changing this name
type Proxy struct {
	Endpoints      Endpoints
	RequestHandler *mux.Router
}

func (p Proxy) handleEndpoints(ctx context.Context) error {
	for _, endpoint := range p.Endpoints.Endpoints {
		if endpoint.UseSharedOptions {
			endpoint.Options = p.Endpoints.SharedOptions
		}

		router, err := newRouter(endpoint.TransportMode)
		if err != nil {
			// TODO: decide if we should wrap this
			return err
		}

		authorizer, err := NewAuthorizer(*endpoint.Options.AuthorizerOptions)
		if err != nil {
			return err
		}

		gateway := Gateway{
			Authorizer: authorizer,
			Router:     router,
			Endpoint:   endpoint,
		}

		gateway.handleEndpoint(ctx, endpoint, p.RequestHandler)
	}

	return nil
}
