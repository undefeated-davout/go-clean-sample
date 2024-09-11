package user

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepo domain.UserRepository
}

func NewCreateUserUseCase(userRepo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) CreateUser(name, email, password string) error {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 新しいユーザーエンティティを作成
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	// リポジトリを介してユーザーを保存
	return uc.userRepo.Save(user)
}
