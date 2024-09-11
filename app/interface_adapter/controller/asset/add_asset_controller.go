package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/asset"
)

type AddAssetController struct {
	addAssetUseCase *asset.AddAssetUseCase
}

func NewAddAssetController(addAssetUC *asset.AddAssetUseCase) *AddAssetController {
	return &AddAssetController{addAssetUseCase: addAssetUC}
}

func (c *AddAssetController) Handle(ctx echo.Context) error {
	type Request struct {
		Ticker string  `json:"ticker"`
		Weight float64 `json:"weight"`
	}

	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := c.addAssetUseCase.AddAsset(req.Ticker, req.Weight)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Asset added successfully"})
}
