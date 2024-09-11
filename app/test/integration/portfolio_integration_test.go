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
	conn, err := db.NewMySQLConnection()
	if err != nil {
		t.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer conn.Close()

	portfolioRepo := repository.NewMySQLPortfolioRepository(conn.DB)
	createPortfolioUC := portfolio.NewCreatePortfolioUseCase(portfolioRepo)
	getPortfolioUC := portfolio.NewGetPortfolioUseCase(portfolioRepo)

	testAssets := []domain.Asset{
		{Ticker: "AAPL", Weight: 0.5},
		{Ticker: "GOOGL", Weight: 0.5},
	}

	err = createPortfolioUC.CreatePortfolio("Test Portfolio", testAssets)
	assert.NoError(t, err, "Error should not have occurred during portfolio creation")

	result, err := getPortfolioUC.GetPortfolio(1)
	assert.NoError(t, err, "Error should not have occurred during portfolio retrieval")
	assert.Equal(t, "Test Portfolio", result.Name)
	assert.Equal(t, 2, len(result.Assets), "The portfolio should have two assets")
}
