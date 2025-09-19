package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/o5h/quiz/pkg/services"
)

type AuthController interface {
	Login(ctx echo.Context) error
}

type authController struct {
	srv services.AuthService
}

func NewAuthController(srv services.AuthService) AuthController {
	return &authController{
		srv: srv,
	}
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (c *authController) Login(ctx echo.Context) error {
	var req LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{"error": "invalid request"})
	}

	token, err := c.srv.Login(req.Username, req.Password)
	if err != nil {
		return ctx.JSON(401, ErrorResponse{Reason: err.Error(), Message: "authentication failed"})
	}

	return ctx.JSON(200, &LoginResponse{AccessToken: token})
}
