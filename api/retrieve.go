package api

import (
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
)

func ConfigRetrieveRouter(e *echo.Group, node *core.LightNode) {

	content := e.Group("/retrieve")
	content.POST("/cid", func(c echo.Context) error {
		return nil
	})

}
