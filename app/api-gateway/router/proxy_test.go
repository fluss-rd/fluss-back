package router

import (
	"context"
	"net/http"
	"testing"

	authRepoMock "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth/mocks"
)

func TestProxyHandleRoutes(t *testing.T) {
	endpoints := Endpoints{
		Endpoints: []Endpoint{
			{
				Path:             "/account/login",
				RemotePath:       "/login",
				RemotHost:        "http://accounts:5000",
				Method:           http.MethodPost,
				Authorized:       false,
				UseSharedOptions: true,
				TransportMode:    TransportModeHTTP,
			},
		},
	}

	repo := authRepoMock.Repository{}

	proxy := Proxy{
		Endpoints: []Endpoints{endpoints},
	}

	proxy.HandleEndpoints(context.Background(), &repo)

}
