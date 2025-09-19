package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/o5h/config"
	"github.com/o5h/quiz/pkg/context"
	"github.com/o5h/quiz/pkg/services"
)

type ControllerStatus struct {
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

func Start() {
	server := echo.New()

	authCtrl := NewAuthController(services.NewAuthService())
	authGroup := server.Group("/auth")
	{ // Auth routes
		authGroup.POST("/login", authCtrl.Login)
	}

	server.GET("/health", func(c echo.Context) error {
		return c.JSON(200, &ControllerStatus{
			Uptime:  "TODO",
			Version: "1.0.0",
		})
	})

	server.POST("/shutdown", func(c echo.Context) error {
		c.JSON(200, "Shutting down...")
		context.Shutdown()
		return nil
	})

	// Start server
	server.Logger.Fatal(server.Start(config.Get("server.address", ":8080")))
}
