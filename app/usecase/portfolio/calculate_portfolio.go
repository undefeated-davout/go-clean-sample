package portfolio

import (
	"errors"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type CalculatePortfolioUseCase struct {
	portfolioRepo domain.PortfolioRepository
}

type PortfolioCalculator interface {
	CalculatePortfolio(id int) (expectedReturn, risk, sharpeRatio float64, err error)
}

var _ PortfolioCalculator = (*CalculatePortfolioUseCase)(nil)

func NewCalculatePortfolioUseCase(portfolioRepo domain.PortfolioRepository) *CalculatePortfolioUseCase {
	return &CalculatePortfolioUseCase{portfolioRepo: portfolioRepo}
}

// ポートフォリオのリスクやリターンを計算するロジック
func (uc *CalculatePortfolioUseCase) CalculatePortfolio(id int) (expectedReturn, risk, sharpeRatio float64, err error) {
	portfolio, err := uc.portfolioRepo.GetByID(id)
	if err != nil {
		return 0, 0, 0, err
	}

	if len(portfolio.Assets) == 0 {
		return 0, 0, 0, errors.New("no assets in the portfolio")
	}

	// TODO: ここにリスクやリターンを計算するロジックを追加する
	// 期待リターン、リスク、シャープレシオを返す
	return 0.08, 0.15, 1.2, nil
}
