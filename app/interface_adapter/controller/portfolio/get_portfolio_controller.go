package portfolio

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
)

type GetPortfolioController struct {
	getPortfolioUseCase *portfolio.GetPortfolioUseCase
}

func NewGetPortfolioController(getPortfolioUC *portfolio.GetPortfolioUseCase) *GetPortfolioController {
	return &GetPortfolioController{getPortfolioUseCase: getPortfolioUC}
}

func (c *GetPortfolioController) Handle(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid portfolio ID"})
	}

	portfolio, err := c.getPortfolioUseCase.GetPortfolio(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, portfolio)
}
