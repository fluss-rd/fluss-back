package service

import (
	"context"
	"errors"

	"github.com/flussrd/fluss-back/app/accounts/models"
	rolesRepository "github.com/flussrd/fluss-back/app/accounts/repositories/roles"
	usersRepository "github.com/flussrd/fluss-back/app/accounts/repositories/users"
	"github.com/flussrd/fluss-back/app/accounts/shared/httputils"
	"github.com/flussrd/fluss-back/app/accounts/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	userIDPrefix = "USR"
)

var (
	// ErrMissingEmail missing email
	ErrMissingEmail = httputils.NewBadRequestError("missing email")
	// ErrMissingPassword missing password
	ErrMissingPassword = httputils.NewBadRequestError("missing password")
	// ErrMissingName missing name
	ErrMisssingName = httputils.NewBadRequestError("missing name")
	// ErrMissingRole missing role
	ErrMissingRole = httputils.NewBadRequestError("missing role")
	// ErrPasswordHashingFailed password hashing failed
	ErrPasswordHashingFailed = errors.New("hashing password failed")
	// ErrGeneratingIDFailed generating id failed
	ErrGeneratingIDFailed = errors.New("generating id failed")
)

var (
	generatePasswordHashFunction func(password []byte, cost int) ([]byte, error)
	generateIDFunction           func(prefix string) (string, error)
)

func init() {
	generatePasswordHashFunction = bcrypt.GenerateFromPassword
	generateIDFunction = utils.GenerateID
}

type service struct {
	usersRepo usersRepository.Repository
	rolesRepo rolesRepository.Repository
}

// NewService returns a new service entity to be able to execuse business logic
func NewService(usersRepo usersRepository.Repository, rolesRepo rolesRepository.Repository) Service {
	return service{
		usersRepo: usersRepo,
		rolesRepo: rolesRepo,
	}
}

// CreateUser creates a new user
func (s service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	err := validateCreateUserParams(user)
	if err != nil {
		return models.User{}, err
	}

	hashedPassword, err := generatePasswordHashFunction([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, ErrPasswordHashingFailed
	}

	user.Password = string(hashedPassword)

	id, err := generateIDFunction(userIDPrefix)
	if err != nil {
		return models.User{}, ErrGeneratingIDFailed
	}

	user.UserID = id

	insertedUser, err := s.usersRepo.SaveUser(ctx, user)
	if errors.As(err, &usersRepository.ErrDuplicateFields{}) {
		return models.User{}, httputils.NewBadRequestError(err.Error())
	}

	return insertedUser, nil
}

func validateCreateUserParams(user models.User) error {
	if user.Email == "" {
		return ErrMissingEmail
	}

	if user.Password == "" {
		return ErrMissingPassword
	}

	if user.Name == "" {
		return ErrMisssingName
	}

	if user.RoleName == "" {
		return ErrMissingRole
	}

	return nil
}

// AddRoleToUser adds a role to a user
func (s service) AddRoleToUser(ctx context.Context, roleName string, userID string) error {
	return nil
}

// CreateRole creates a new role
func (s service) CreateRole(ctx context.Context, role models.Role) error {
	return s.rolesRepo.CreateRole(ctx, role)
}

// UpdateRole updates a role
func (s service) UpdateRole(ctx context.Context, role models.Role) error {
	return nil
}

// GetRoles returns all the roles
func (s service) GetRoles(ctx context.Context) ([]models.Role, error) {
	// TODO: handle repo specific errors
	return s.rolesRepo.GetRoles(ctx)
}
