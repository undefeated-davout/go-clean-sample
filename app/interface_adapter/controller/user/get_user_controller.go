package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/user"
)

type GetUserController struct {
	getUserUseCase *user.GetUserUseCase
}

func NewGetUserController(getUserUC *user.GetUserUseCase) *GetUserController {
	return &GetUserController{getUserUseCase: getUserUC}
}

func (c *GetUserController) Handle(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	if email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required"})
	}

	user, err := c.getUserUseCase.GetUserByEmail(email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}
