package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/o5h/config"
)

type ControllerStatus struct {
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

func Start() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, &ControllerStatus{
			Uptime:  "TODO",
			Version: "1.0.0",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(config.Get("server.address", ":8080")))
}
