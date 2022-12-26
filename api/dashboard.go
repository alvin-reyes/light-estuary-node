package api

import (
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
)

func ConfigureDashboardRouter(e *echo.Group, node *core.LightNode) {

	gatewayHandler.node = node.Node
	e.GET("/dashboard", nil)
}
