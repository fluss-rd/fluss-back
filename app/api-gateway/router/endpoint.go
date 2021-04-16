package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
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

// Router defines the methods for generating routes
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

		if endpoint.Authorized {
			r.handler.Handle(endpoint.Path, authMiddleware(ctx, endpoint, handleRequest(endpoint, remoteURL), r.authRepo)).Methods(endpoint.Method)
			continue
		}

		r.handler.Handle(endpoint.Path, handleRequest(endpoint, remoteURL)).Methods(endpoint.Method)
	}

	return nil
}

func modifyRequest(r *http.Request, remoteURL *url.URL) {
	// Update the headers to allow for SSL redirection
	r.URL.Host = remoteURL.Host
	r.URL.Scheme = remoteURL.Scheme
	r.URL.Path = remoteURL.Path // we need this to come from the endpoint to not have
	r.Header.Set("X-Forwarded-Host", r.Host)
	r.Host = remoteURL.Host
}

func getProxy(url *url.URL) (*httputil.ReverseProxy, error) {
	return httputil.NewSingleHostReverseProxy(url), nil
}

func getActualURL(endpoint Endpoint, vars map[string]string) (*url.URL, error) {
	wholeRemoteEndpoint := endpoint.RemotHost + endpoint.RemotePath

	for k, v := range vars {
		wholeRemoteEndpoint = strings.ReplaceAll(wholeRemoteEndpoint, fmt.Sprintf("{%s}", k), v)
	}

	url, err := url.Parse(wholeRemoteEndpoint)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func handleRequest(endpoint Endpoint, remoteURL *url.URL) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		remoteURL, err := getActualURL(endpoint, mux.Vars(r))
		if err != nil {
			fmt.Println("failed to get url: " + err.Error())
			httputils.RespondInternalServerError(rw)
			return
		}

		remoteHostURL, err := url.Parse(endpoint.RemotHost)
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

func authMiddleware(ctx context.Context, endpoint Endpoint, handler http.HandlerFunc, repo authRepo.Repository) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// token, err := getToken(*r)
		// if err != nil {
		// 	httputils.RespondWithError(rw, err)
		// 	return
		// }

		// resource, err := getResourceFromEndpoint(endpoint)
		// if err != nil {
		// 	httputils.RespondWithError(rw, err)
		// 	return
		// }

		// action, err := getActionFromMethod(r.Method)
		// if err != nil {
		// 	httputils.RespondWithError(rw, err)
		// 	return
		// }

		// authorizer := authorization.NewAuthorizer(repo, jwt.SigningMethodHS256)

		// // TODO: handle error, this can be a reason to return 500
		// isAuthorized, _ := authorizer.Validate(ctx, token, resource, action)
		// if !isAuthorized {
		// 	// TODO: make this beautiful
		// 	httputils.RespondWithError(rw, httputils.ForbiddenError)

		// 	return
		// }

		handler.ServeHTTP(rw, r)
	}
}

func getResourceFromEndpoint(endpoint Endpoint) (string, error) {
	splitted := strings.Split(endpoint.Path, "/")
	// TODO: handle when this is of unexpected length
	// If the last part of the path contais a {, that means is a variable like : modules/{id}, so we should return the one before that
	if strings.Contains(splitted[len(splitted)-1], "{") {
		return splitted[len(splitted)-2], nil
	}

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
