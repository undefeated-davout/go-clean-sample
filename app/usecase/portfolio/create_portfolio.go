package portfolio

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type CreatePortfolioUseCase struct {
	portfolioRepo domain.PortfolioRepository
}

func NewCreatePortfolioUseCase(portfolioRepo domain.PortfolioRepository) *CreatePortfolioUseCase {
	return &CreatePortfolioUseCase{portfolioRepo: portfolioRepo}
}

func (uc *CreatePortfolioUseCase) CreatePortfolio(name string, assets []domain.Asset) error {
	portfolio := domain.Portfolio{
		Name:   name,
		Assets: assets,
	}

	// リポジトリを介してポートフォリオを保存
	return uc.portfolioRepo.Save(portfolio)
}
