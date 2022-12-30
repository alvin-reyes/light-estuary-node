package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// This stores allow the consumer to get all the output of this light node in
// request-driven way
func ConfigOutputRouter(e *echo.Group) {
	e.GET("/output", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Ok")
	})

	// output all the data from the database
	e.GET("/output-all", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Ok")
	})
}
