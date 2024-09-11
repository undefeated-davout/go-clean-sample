package portfolio

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
)

type CalculatePortfolioController struct {
	calculatePortfolioUseCase *portfolio.CalculatePortfolioUseCase
}

func NewCalculatePortfolioController(calculatePortfolioUC *portfolio.CalculatePortfolioUseCase) *CalculatePortfolioController {
	return &CalculatePortfolioController{calculatePortfolioUseCase: calculatePortfolioUC}
}

func (c *CalculatePortfolioController) Handle(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid portfolio ID"})
	}

	expectedReturn, risk, sharpeRatio, err := c.calculatePortfolioUseCase.CalculatePortfolio(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]float64{
		"expectedReturn": expectedReturn,
		"risk":           risk,
		"sharpeRatio":    sharpeRatio,
	})
}
