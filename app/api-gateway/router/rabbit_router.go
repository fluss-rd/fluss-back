package router

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/shared/rabbit"
)

type rabbitRouter struct {
	endpoint     Endpoint
	rabbitClient rabbit.RabbitClient
}

func newRabbitMqRouter(endpoint Endpoint, rabbitClient rabbit.RabbitClient) Router {
	return rabbitRouter{
		endpoint:     endpoint,
		rabbitClient: rabbitClient,
	}
}

func (router rabbitRouter) Route() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			httputils.RespondInternalServerError(rw)
			return
		}

		err = router.rabbitClient.Publish(context.Background(), router.endpoint.ExchangeName, router.endpoint.RoutingKey, requestBody)
		if err != nil {
			httputils.RespondInternalServerError(rw)
			return
		}

		httputils.RespondJSON(rw, http.StatusAccepted, "")
	}
}
