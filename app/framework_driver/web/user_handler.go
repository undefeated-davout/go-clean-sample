package web

import (
	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/user"
)

type UserHandler struct {
	createUserController      *user.CreateUserController
	authenticateUserController *user.AuthenticateUserController
	getUserController         *user.GetUserController
}

func NewUserHandler(createUserController *user.CreateUserController, authenticateUserController *user.AuthenticateUserController, getUserController *user.GetUserController) *UserHandler {
	return &UserHandler{
		createUserController:      createUserController,
		authenticateUserController: authenticateUserController,
		getUserController:         getUserController,
	}
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/users", h.createUserController.Handle)
	e.POST("/users/auth", h.authenticateUserController.Handle)
	e.GET("/users", h.getUserController.Handle)
}
