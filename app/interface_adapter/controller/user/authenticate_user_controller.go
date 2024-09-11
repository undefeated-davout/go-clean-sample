package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/user"
)

type AuthenticateUserController struct {
	authenticateUserUseCase *user.AuthenticateUserUseCase
}

func NewAuthenticateUserController(authenticateUserUC *user.AuthenticateUserUseCase) *AuthenticateUserController {
	return &AuthenticateUserController{authenticateUserUseCase: authenticateUserUC}
}

func (c *AuthenticateUserController) Handle(ctx echo.Context) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user, err := c.authenticateUserUseCase.Authenticate(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Authentication failed"})
	}

	return ctx.JSON(http.StatusOK, user)
}
