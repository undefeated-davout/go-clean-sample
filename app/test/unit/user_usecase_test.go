package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/test/mock"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/user"
)

func TestCreateUser(t *testing.T) {
	mockRepo := mock.NewMockUserRepository()
	createUserUC := user.NewCreateUserUseCase(mockRepo)

	err := createUserUC.CreateUser("John Doe", "john@example.com", "password123")
	assert.NoError(t, err, "Error should not have occurred during user creation")

	result, err := mockRepo.GetByEmail("john@example.com")
	assert.NoError(t, err, "User should have been found")
	assert.Equal(t, "John Doe", result.Name)
}

func TestAuthenticateUser(t *testing.T) {
	mockRepo := mock.NewMockUserRepository()
	authenticateUserUC := user.NewAuthenticateUserUseCase(mockRepo)

	mockRepo.Save(domain.User{Name: "John Doe", Email: "john@example.com", Password: "password123"})

	user, err := authenticateUserUC.Authenticate("john@example.com", "password123")
	assert.NoError(t, err, "User should be authenticated successfully")
	assert.Equal(t, "John Doe", user.Name)
}
