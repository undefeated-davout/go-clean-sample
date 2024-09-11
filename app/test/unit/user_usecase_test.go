package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/test/mock"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/user"
)

func TestCreateUser(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockUserRepository()
	createUserUC := user.NewCreateUserUseCase(mockRepo)

	// ユーザー作成のテスト
	err := createUserUC.CreateUser("John Doe", "john@example.com", "password123")
	assert.NoError(t, err, "Error should not have occurred during user creation")

	// ユーザーが正しく作成されたか確認
	result, err := mockRepo.GetByEmail("john@example.com")
	assert.NoError(t, err, "User should have been found")
	assert.Equal(t, "John Doe", result.Name)
}

func TestAuthenticateUser(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockUserRepository()
	authenticateUserUC := user.NewAuthenticateUserUseCase(mockRepo)

	// 事前にユーザーを追加
	mockRepo.Save(domain.User{Name: "John Doe", Email: "john@example.com", Password: "password123"})

	// 認証のテスト
	user, err := authenticateUserUC.Authenticate("john@example.com", "password123")
	assert.NoError(t, err, "User should be authenticated successfully")
	assert.Equal(t, "John Doe", user.Name)
}
