package api

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"light-estuary-node/core"
)

func ConfigureDashboardRouter(e *echo.Group, node *core.LightNode) {
	gatewayHandler.node = node.Node
	e.GET("/dashboard", ServeDashboard)
}

func ServeDashboard(c echo.Context) error {
	templates, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		return err
	}
	templates.Lookup("dashboard.html")
	templates.Execute(c.Response().Writer, c)

	return nil
}
