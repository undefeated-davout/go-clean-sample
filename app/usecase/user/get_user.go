package user

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type GetUserUseCase struct {
	userRepo domain.UserRepository
}

func NewGetUserUseCase(userRepo domain.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

func (uc *GetUserUseCase) GetUserByID(id int) (*domain.User, error) {
	// リポジトリを使ってユーザーをIDで取得
	return uc.userRepo.GetByID(id)
}

func (uc *GetUserUseCase) GetUserByEmail(email string) (*domain.User, error) {
	// リポジトリを使ってユーザーをメールアドレスで取得
	return uc.userRepo.GetByEmail(email)
}
