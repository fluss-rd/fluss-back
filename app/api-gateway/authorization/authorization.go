package authorization

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flussrd/fluss-back/app/accounts/models"
	repository "github.com/flussrd/fluss-back/app/api-gateway/repositories/auth"
)

var (
	ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
)

type claims struct {
	RoleName string `json:"roleName"`
	Sub      string `json:"sub"`
}

//
type Authorizer interface {
	Validate(ctx context.Context, token, resource, action string) (bool, error)
}

type authorizer struct {
	repo          repository.Repository
	signingMethod jwt.SigningMethod
}

func NewAuthorizer(authRepository repository.Repository, tokenSigningMethod jwt.SigningMethod) Authorizer {
	return authorizer{repo: authRepository, signingMethod: tokenSigningMethod}
}

// Validate validates
func (a authorizer) Validate(ctx context.Context, token string, resource string, action string) (bool, error) {
	authToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t.Method)

		if t.Method.Alg() != a.signingMethod.Alg() {
			return nil, ErrInvalidTokenSigningMethod
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	if !authToken.Valid {
		return false, nil
	}

	claims, err := getTokenClaims(authToken)
	if err != nil {
		return false, err
	}

	// TODO: validate claims has the requried fields
	role, err := a.getRole(ctx, claims)
	if err != nil {
		return false, err
	}

	return checkPermissions(role, resource, action), nil
}

func checkPermissions(role models.Role, resource string, action string) bool {
	for _, permission := range role.Permissions {
		if (string(permission.Resource) == resource || permission.Resource == "*") && (string(permission.Action) == action || permission.Action == "*") {
			return true
		}
	}

	return false
}

func (a authorizer) getRole(ctx context.Context, claims claims) (models.Role, error) {
	// TODO: handle when no roles
	return a.repo.GetRole(ctx, claims.RoleName)
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
