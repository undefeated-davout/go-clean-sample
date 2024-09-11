package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/test/mock"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
)

func TestCreatePortfolio(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockPortfolioRepository()
	createPortfolioUC := portfolio.NewCreatePortfolioUseCase(mockRepo)

	// ポートフォリオ作成のテスト
	testAssets := []domain.Asset{
		{Ticker: "AAPL", Weight: 0.5},
		{Ticker: "GOOGL", Weight: 0.5},
	}

	err := createPortfolioUC.CreatePortfolio("Test Portfolio", testAssets)
	assert.NoError(t, err, "Error should not have occurred during portfolio creation")
}

func TestGetPortfolio(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockPortfolioRepository()
	getPortfolioUC := portfolio.NewGetPortfolioUseCase(mockRepo)

	// ポートフォリオを事前に保存
	mockRepo.Save(domain.Portfolio{
		ID:   1,
		Name: "Test Portfolio",
		Assets: []domain.Asset{
			{Ticker: "AAPL", Weight: 0.5},
			{Ticker: "GOOGL", Weight: 0.5},
		},
	})

	// ポートフォリオ取得のテスト
	result, err := getPortfolioUC.GetPortfolio(1)
	assert.NoError(t, err, "Error should not have occurred during portfolio retrieval")
	assert.Equal(t, "Test Portfolio", result.Name)
	assert.Equal(t, 2, len(result.Assets))
}
