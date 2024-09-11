package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/asset"
)

type RemoveAssetController struct {
	removeAssetUseCase *asset.RemoveAssetUseCase
}

func NewRemoveAssetController(removeAssetUC *asset.RemoveAssetUseCase) *RemoveAssetController {
	return &RemoveAssetController{removeAssetUseCase: removeAssetUC}
}

func (c *RemoveAssetController) Handle(ctx echo.Context) error {
	ticker := ctx.Param("ticker")
	if ticker == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Ticker is required"})
	}

	err := c.removeAssetUseCase.RemoveAsset(ticker)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Asset removed successfully"})
}
