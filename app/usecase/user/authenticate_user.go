package user

import (
	"errors"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUserUseCase struct {
	userRepo domain.UserRepository
}

type UserAuthenticator interface {
	Authenticate(email, password string) (*domain.User, error)
}

var _ UserAuthenticator = (*AuthenticateUserUseCase)(nil)

func NewAuthenticateUserUseCase(userRepo domain.UserRepository) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{userRepo: userRepo}
}

func (uc *AuthenticateUserUseCase) Authenticate(email, password string) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
