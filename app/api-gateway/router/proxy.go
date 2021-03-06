package router

import (
	"context"

	authRepo "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
	"github.com/gorilla/mux"
)

// TODO: consider changing this name
type Proxy struct {
	Endpoints      []Endpoints
	RequestHandler *mux.Router
}

func (p Proxy) HandleEndpoints(ctx context.Context, repo authRepo.Repository, rabbitclient rabbit.RabbitClient) error {
	for _, endpoints := range p.Endpoints {
		for _, endpoint := range endpoints.Endpoints {
			if endpoint.UseSharedOptions {
				endpoint.Options = endpoints.SharedOptions
			}

			router, err := newRouter(endpoint, endpoint.TransportMode, rabbitclient)
			if err != nil {
				// TODO: decide if we should wrap this
				return err
			}

			gateway := Gateway{
				Router:   router,
				Endpoint: endpoint,
			}

			if endpoint.Authorized {
				// TODO: validate that the options are not nil
				authorizer, err := NewAuthorizer(*endpoint.Options.AuthorizerOptions, repo)
				if err != nil {
					return err
				}

				gateway.Authorizer = authorizer
			}

			gateway.handleEndpoint(ctx, endpoint, p.RequestHandler)
		}

	}

	return nil
}
