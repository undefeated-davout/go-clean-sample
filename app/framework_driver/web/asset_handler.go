package web

import (
	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/asset"
)

type AssetHandler struct {
	addAssetController    *asset.AddAssetController
	removeAssetController *asset.RemoveAssetController
}

func NewAssetHandler(addAssetController *asset.AddAssetController, removeAssetController *asset.RemoveAssetController) *AssetHandler {
	return &AssetHandler{
		addAssetController:    addAssetController,
		removeAssetController: removeAssetController,
	}
}

func (h *AssetHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/assets", h.addAssetController.Handle)
	e.DELETE("/assets/:ticker", h.removeAssetController.Handle)
}
