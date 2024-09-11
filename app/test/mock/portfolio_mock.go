package mock

import (
	"errors"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MockPortfolioRepository struct {
	portfolios map[int]*domain.Portfolio
	nextID     int
}

func NewMockPortfolioRepository() *MockPortfolioRepository {
	return &MockPortfolioRepository{
		portfolios: make(map[int]*domain.Portfolio),
		nextID:     1,
	}
}

func (r *MockPortfolioRepository) Save(portfolio domain.Portfolio) error {
	portfolio.ID = r.nextID
	r.portfolios[r.nextID] = &portfolio
	r.nextID++
	return nil
}

func (r *MockPortfolioRepository) GetByID(id int) (*domain.Portfolio, error) {
	portfolio, exists := r.portfolios[id]
	if !exists {
		return nil, errors.New("portfolio not found")
	}
	return portfolio, nil
}

func (r *MockPortfolioRepository) Delete(id int) error {
	delete(r.portfolios, id)
	return nil
}
