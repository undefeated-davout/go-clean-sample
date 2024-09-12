package mock

import (
	"errors"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	users map[string]*domain.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *MockUserRepository) Save(user domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	r.users[user.Email] = &user
	return nil
}

func (r *MockUserRepository) GetByID(id int) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MockUserRepository) Delete(id int) error {
	for email, user := range r.users {
		if user.ID == id {
			delete(r.users, email)
			return nil
		}
	}
	return errors.New("user not found")
}
