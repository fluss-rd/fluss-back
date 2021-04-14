package service

import (
	"context"
	"testing"

	"github.com/flussrd/fluss-back/app/accounts/models"
	rolesRepoMock "github.com/flussrd/fluss-back/app/accounts/repositories/roles/mock"
	usersRepoMock "github.com/flussrd/fluss-back/app/accounts/repositories/users/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	c := require.New(t)

	usersRepo := usersRepoMock.Repository{}
	rolesRepo := rolesRepoMock.Repository{}

	service := NewService(&usersRepo, &rolesRepo)

	user := models.User{
		PhoneNumber: "+1809000000",
		Email:       "email@email.com",
		Password:    "nvjkfe",
		Name:        "Francia",
		RoleName:    "admin",
	}

	// User that is sent to the repo
	trasformedUser := models.User{
		PhoneNumber: "+1809000000",
		Email:       "email@email.com",
		Password:    "123",
		Name:        "Francia",
		RoleName:    "admin",
		UserID:      "USR123",
	}

	ctx := context.Background()

	generateIDFunction = func(prefix string) (string, error) {
		return "USR123", nil
	}

	generatePasswordHashFunction = func(password []byte, cost int) ([]byte, error) {
		return []byte("123"), nil
	}

	usersRepo.On("SaveUser", ctx, trasformedUser).Return(trasformedUser, nil)

	insertedUser, err := service.CreateUser(ctx, user)
	c.Nil(err)
	c.NotNil(insertedUser)
	c.NotEmpty(insertedUser.UserID)
}

func TestAddRoleToUser(t *testing.T) {
	c := require.New(t)

	usersRepo := usersRepoMock.Repository{}
	rolesRepo := rolesRepoMock.Repository{}

	service := NewService(&usersRepo, &rolesRepo)

	err := service.AddRoleToUser(context.Background(), "admin", "USR123")
	c.Nil(err)
}

func TestCreaterole(t *testing.T) {
	c := require.New(t)

	usersRepo := usersRepoMock.Repository{}
	rolesRepo := rolesRepoMock.Repository{}

	ctx := context.Background()

	role := models.Role{
		Name: "admin",
		Permissions: []models.Permission{
			{
				Resource: "modules",
				Action:   "*",
			},
		},
	}

	rolesRepo.On("CreateRole", ctx, role).Return(nil)

	service := NewService(&usersRepo, &rolesRepo)

	err := service.CreateRole(ctx, role)

	c.Nil(err)
}

func TestUpdateRole(t *testing.T) {
	c := require.New(t)

	usersRepo := usersRepoMock.Repository{}
	rolesRepo := rolesRepoMock.Repository{}

	service := NewService(&usersRepo, &rolesRepo)

	err := service.UpdateRole(context.Background(), models.Role{
		Name: "admin",
		Permissions: []models.Permission{
			{
				Resource: "modules",
				Action:   "*",
			},
		},
	})

	c.Nil(err)
}
