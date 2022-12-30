package api

import (
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
)

func ConfigDealsRouter(e *echo.Group, node *core.LightNode) {
	deals := e.Group("/deals")
	deals.GET("/list", func(c echo.Context) error {
		return c.JSON(200, "Ok")
	})
}
