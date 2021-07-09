package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/flussrd/fluss-back/app/accounts/models"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	authRepo "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
)

type AuthorizerType string

const (
	AuthorizerTypeJWT AuthorizerType = "jwt"
)

var (
	// ErrInvalidTokenSigningMethod invalid token signing method
	// ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
	// ErrMissingRoleName missing role name
	ErrMissingRoleName = errors.New("missing role name")
	// ErrMissingSub missing sub
	ErrMissingSub = errors.New("missing sub")
)

type AuthorizeResult string

// authentication and autorization
type Authorizer interface {
	Authorize(ctx context.Context, req *http.Request, resource string, action string) (string, error)
}

func NewAuthorizer(options AuthorizerOptions, repo authRepo.Repository) (Authorizer, error) {
	switch options.AuthType {
	case AuthorizerTypeJWT:
		return newJWTAuthorizer(options, repo), nil
	}

	return nil, errors.New("invalid authorizer")
}

type claims struct {
	RoleName string `json:"roleName"`
	Sub      string `json:"sub"`
}

type jwtAuthorizer struct {
	repo              authRepo.Repository
	authorizerOptions AuthorizerOptions
}

func newJWTAuthorizer(options AuthorizerOptions, repo authRepo.Repository) Authorizer {
	return jwtAuthorizer{
		authorizerOptions: options,
		repo:              repo,
	}
}

func (authorizer jwtAuthorizer) Authorize(ctx context.Context, req *http.Request, resource string, action string) (string, error) {
	token, err := authorizer.getToken(req)
	if err != nil {
		return "", err
	}

	isAuthorized, sub, _ := authorizer.validateToken(ctx, token, resource, action)
	if !isAuthorized {
		return "", httputils.ForbiddenError
	}

	return sub, nil
}

func (authorizer jwtAuthorizer) getToken(r *http.Request) (string, error) {
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

func (authorizer jwtAuthorizer) validateToken(ctx context.Context, token string, resource, action string) (bool, string, error) {
	authToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != authorizer.authorizerOptions.JWTSigningMethod.Alg() {
			return nil, ErrInvalidTokenSigningMethod
		}

		return []byte(authorizer.authorizerOptions.JwtSigningSecret), nil
	})

	if err != nil {
		return false, "", err
	}

	if !authToken.Valid {
		return false, "", nil
	}

	tokenClaims, err := getTokenClaims(authToken)
	if err != nil {
		return false, "", err
	}

	err = validateTokenClaims(tokenClaims)
	if err != nil {
		return false, "", err
	}

	role, err := authorizer.repo.GetRole(ctx, tokenClaims.RoleName)
	if err != nil {
		return false, "", err
	}

	return checkPermissions(role, resource, action), tokenClaims.Sub, nil
}

func validateTokenClaims(claims claims) error {
	if claims.RoleName == "" {
		return ErrMissingRoleName
	}

	if claims.Sub == "" {
		return ErrMissingSub
	}

	return nil
}

func getTokenClaims(token *jwt.Token) (claims, error) {
	claimsBytes, err := json.Marshal(token.Claims)
	if err != nil {
		return claims{}, err
	}

	tokenClaims := claims{}
	err = json.Unmarshal(claimsBytes, &tokenClaims)
	if err != nil {
		return claims{}, err
	}

	return tokenClaims, nil
}

func checkPermissions(role models.Role, resource string, desiredAction string) bool {
	for _, permission := range role.Permissions {
		for _, action := range permission.Actions {
			if (string(permission.Resource) == resource || permission.Resource == "*") && (string(action) == desiredAction || action == "*") {
				return true
			}
		}
	}

	return false
}
