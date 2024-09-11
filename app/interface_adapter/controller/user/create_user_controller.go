package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/undefeated-davout/portfolio-simulator/app/usecase/user"
)

type CreateUserController struct {
	createUserUseCase *user.CreateUserUseCase
}

func NewCreateUserController(createUserUC *user.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{createUserUseCase: createUserUC}
}

func (c *CreateUserController) Handle(ctx echo.Context) error {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(Request)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := c.createUserUseCase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}
