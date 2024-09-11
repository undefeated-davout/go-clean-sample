package portfolio

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type GetPortfolioUseCase struct {
	portfolioRepo domain.PortfolioRepository
}

func NewGetPortfolioUseCase(portfolioRepo domain.PortfolioRepository) *GetPortfolioUseCase {
	return &GetPortfolioUseCase{portfolioRepo: portfolioRepo}
}

func (uc *GetPortfolioUseCase) GetPortfolio(id int) (*domain.Portfolio, error) {
	return uc.portfolioRepo.GetByID(id)
}
