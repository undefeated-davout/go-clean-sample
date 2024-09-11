package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/framework_driver/db"
	"github.com/undefeated-davout/portfolio-simulator/app/framework_driver/db/repository"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
)

func TestPortfolioIntegration(t *testing.T) {
	// MySQL接続をセットアップ
	conn, err := db.NewMySQLConnection()
	if err != nil {
		t.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer conn.Close()

	// ポートフォリオリポジトリの準備
	portfolioRepo := repository.NewMySQLPortfolioRepository(conn.DB)
	createPortfolioUC := portfolio.NewCreatePortfolioUseCase(portfolioRepo)
	getPortfolioUC := portfolio.NewGetPortfolioUseCase(portfolioRepo)

	// テスト用のポートフォリオを作成
	testAssets := []domain.Asset{
		{Ticker: "AAPL", Weight: 0.5},
		{Ticker: "GOOGL", Weight: 0.5},
	}

	err = createPortfolioUC.CreatePortfolio("Test Portfolio", testAssets)
	assert.NoError(t, err, "Error should not have occurred during portfolio creation")

	// ポートフォリオの取得
	result, err := getPortfolioUC.GetPortfolio(1) // ID 1のポートフォリオを取得（仮のID）
	assert.NoError(t, err, "Error should not have occurred during portfolio retrieval")
	assert.Equal(t, "Test Portfolio", result.Name)
	assert.Equal(t, 2, len(result.Assets), "The portfolio should have two assets")
}
