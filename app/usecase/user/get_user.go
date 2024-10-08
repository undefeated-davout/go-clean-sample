package user

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type GetUserUseCase struct {
	userRepo domain.UserRepository
}

type UserGetter interface {
	GetUserByID(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

var _ UserGetter = (*GetUserUseCase)(nil)

func NewGetUserUseCase(userRepo domain.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

func (uc *GetUserUseCase) GetUserByID(id int) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *GetUserUseCase) GetUserByEmail(email string) (*domain.User, error) {
	return uc.userRepo.GetByEmail(email)
}
