package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/api-gateway/authorization"
	authRepo "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
	"github.com/gorilla/mux"
)

var (
	// ErrParseURLFailed parsing url failed
	ErrParseURLFailed = errors.New("parsing url failed")
	// ErrInvalidTokenSigningMethod invalid token signing method
	ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
	// ErrMissingToken missing authentication token
	ErrMissingToken = httputils.NewBadRequestError("missing authentication token")
)

var (
	// TODO: modify this to have the ActionType type and not string
	methodActions = map[string]string{
		http.MethodGet:    "read",
		http.MethodPost:   "create",
		http.MethodPatch:  "update",
		http.MethodDelete: "delete",
	}
)

// Endpoint defines an endpoint to be routed
type Endpoint struct {
	Path       string
	RemotePath string
	RemotHost  string
	Method     string
	Authorized bool
}

type Router interface {
	generateRoutes(ctx context.Context) error
}

type router struct {
	authRepo  authRepo.Repository
	endpoints []Endpoint
	handler   *mux.Router // TODO: REMOVE THIS DEPENDENCY, this is too coupled
}

// NewRouter returns a new router entity for routing requests
func NewRouter(ctx context.Context, endpoints []Endpoint, repo authRepo.Repository, handler *mux.Router) (Router, error) {
	router := &router{
		endpoints: endpoints,
		authRepo:  repo,
		handler:   handler,
	}

	err := router.generateRoutes(ctx)
	if err != nil {
		return nil, err
	}

	return router, nil
}

func (r *router) generateRoutes(ctx context.Context) error {
	for _, endpoint := range r.endpoints {
		remoteURL, err := url.Parse(endpoint.RemotHost + endpoint.RemotePath)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrParseURLFailed, err.Error())
		}

		remoteHostURL, err := url.Parse(endpoint.RemotHost)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrParseURLFailed, err.Error())
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteHostURL)

		if endpoint.Authorized {
			r.handler.Handle(endpoint.Path, authMiddleware(ctx, handleRequest(proxy, remoteURL), r.authRepo)).Methods(endpoint.Method)
			continue
		}

		r.handler.Handle(endpoint.Path, handleRequest(proxy, remoteURL)).Methods(endpoint.Method)
	}

	return nil
}

func modifyRequest(r *http.Request, remoteURL *url.URL) {
	// Update the headers to allow for SSL redirection
	fmt.Println(remoteURL)
	r.URL.Host = remoteURL.Host
	r.URL.Scheme = remoteURL.Scheme
	r.URL.Path = remoteURL.Path
	r.Header.Set("X-Forwarded-Host", r.Host)
	r.Host = remoteURL.Host
}

func handleRequest(proxy *httputil.ReverseProxy, remoteURL *url.URL) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		modifyRequest(r, remoteURL)
		proxy.ServeHTTP(rw, r)
	}
}

func getToken(r http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", httputils.UnauthorizedError
	}

	splittedToken := strings.Split(authHeader, "Bearer ")
	if len(splittedToken) < 2 {
		return "", httputils.UnauthorizedError
	}

	return splittedToken[1], nil
}

func authMiddleware(ctx context.Context, handler http.HandlerFunc, repo authRepo.Repository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		token, err := getToken(*r)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		resource, err := getResourceFromURL(*r.URL)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		action, err := getActionFromMethod(r.Method)
		if err != nil {
			httputils.RespondWithError(rw, err)
			return
		}

		authorizer := authorization.NewAuthorizer(repo, jwt.SigningMethodHS256)

		// TODO: handle error, this can be a reason to return 500
		isAuthorized, _ := authorizer.Validate(ctx, token, resource, action)
		if !isAuthorized {
			// TODO: make this beautiful
			httputils.RespondWithError(rw, httputils.ForbiddenError)

			return
		}

		handler.ServeHTTP(rw, r)
	}
}

func getResourceFromURL(url url.URL) (string, error) {
	splitted := strings.Split(url.String(), "/")
	// TODO: handle when this is of unexpected length

	return splitted[len(splitted)-1], nil
}

func getActionFromMethod(method string) (string, error) {
	action, ok := methodActions[method]
	if !ok {
		// TODO: define error
		return "", errors.New("method not defined")
	}

	return action, nil
}
