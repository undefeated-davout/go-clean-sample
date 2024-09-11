package web

import (
	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/portfolio"
)

type PortfolioHandler struct {
	createPortfolioController   *portfolio.CreatePortfolioController
	getPortfolioController      *portfolio.GetPortfolioController
	calculatePortfolioController *portfolio.CalculatePortfolioController
}

func NewPortfolioHandler(createPortfolioController *portfolio.CreatePortfolioController, getPortfolioController *portfolio.GetPortfolioController, calculatePortfolioController *portfolio.CalculatePortfolioController) *PortfolioHandler {
	return &PortfolioHandler{
		createPortfolioController:   createPortfolioController,
		getPortfolioController:      getPortfolioController,
		calculatePortfolioController: calculatePortfolioController,
	}
}

// RegisterRoutes は Portfolio に関連するルートを登録します
func (h *PortfolioHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/portfolios", h.createPortfolioController.Handle)
	e.GET("/portfolios/:id", h.getPortfolioController.Handle)
	e.GET("/portfolios/:id/calculate", h.calculatePortfolioController.Handle)
}
