package user

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepo domain.UserRepository
}

type UserCreator interface {
	CreateUser(name, email, password string) error
}

var _ UserCreator = (*CreateUserUseCase)(nil)

func NewCreateUserUseCase(userRepo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) CreateUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return uc.userRepo.Save(user)
}
