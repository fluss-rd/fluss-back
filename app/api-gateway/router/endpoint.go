package router

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	authRepo "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
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
	Path             string
	RemotePath       string
	TransportMode    TransportMode
	RemotHost        string
	Method           string
	Authorized       bool
	Options          EndpointOptions
	UseSharedOptions bool //TODO: maybe we should think of a way of just using CERTAIN shared options
}

type Endpoints struct {
	Endpoints     []Endpoint
	SharedOptions EndpointOptions
}

type EndpointOptions struct {
	*AuthorizerOptions
}

type AuthorizerOptions struct {
	AuthType         AuthorizerType
	JWTSigningMethod jwt.SigningMethod
	JwtSigningSecret string
	AuthRepo         authRepo.Repository
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

func setupPreflightResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
