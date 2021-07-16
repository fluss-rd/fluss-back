package router

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/gorilla/mux"
)

type httpRouter struct {
	endpoint Endpoint
}

func newHttpRouter(endpoint Endpoint) Router {
	return httpRouter{
		endpoint: endpoint,
	}
}

func (router httpRouter) Route() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		setupPreflightResponse(&rw, r)
		if r.Method == http.MethodOptions {
			return
		}

		remoteURL, err := getActualURL(router.endpoint, mux.Vars(r))
		if err != nil {
			fmt.Println("failed to get url: " + err.Error())
			httputils.RespondInternalServerError(rw)
			return
		}

		remoteHostURL, err := url.Parse(router.endpoint.RemotHost)
		if err != nil {
			fmt.Println("failed to get url: " + err.Error())
			httputils.RespondInternalServerError(rw)
			return
		}

		proxy, err := getProxy(remoteHostURL)
		if err != nil {
			fmt.Println("failed to get proxy: " + err.Error())
			httputils.RespondInternalServerError(rw)
			return
		}

		modifyRequest(r, remoteURL)
		proxy.ServeHTTP(rw, r)
	}
}
