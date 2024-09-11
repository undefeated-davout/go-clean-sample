package portfolio

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
)

type CreatePortfolioController struct {
	createPortfolioUseCase *portfolio.CreatePortfolioUseCase
}

func NewCreatePortfolioController(createPortfolioUC *portfolio.CreatePortfolioUseCase) *CreatePortfolioController {
	return &CreatePortfolioController{createPortfolioUseCase: createPortfolioUC}
}

func (c *CreatePortfolioController) Handle(ctx echo.Context) error {
	type AssetRequest struct {
		Ticker string  `json:"ticker"`
		Weight float64 `json:"weight"`
	}

	type Request struct {
		Name   string        `json:"name"`
		Assets []AssetRequest `json:"assets"`
	}

	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var assets []domain.Asset
	for _, a := range req.Assets {
		assets = append(assets, domain.Asset{
			Ticker: a.Ticker,
			Weight: a.Weight,
		})
	}

	err := c.createPortfolioUseCase.CreatePortfolio(req.Name, assets)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Portfolio created successfully"})
}
