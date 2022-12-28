package api

import (
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
)

func ConfigureNodeRouter(e *echo.Group, node *core.LightNode) {
	nodeRouter := e.Group("/node")
	nodeRouter.GET("/info", func(c echo.Context) error {
		json := map[string]interface{}{
			"peerId":  node.Node.Host.ID(),
			"address": node.Node.Host.Addrs(),
		}
		return c.JSON(200, json)
	})
}
