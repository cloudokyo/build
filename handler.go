package build

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// The handler returns application build metadata
//
// Example
//
//	engine.GET("/v1/metadata", build.Handle)
func Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":    User,
		"time":    Time,
		"number":  Number,
		"source":  Source,
		"commit":  Commit,
		"version": Version,
	})
}

// The handler returns request headers metadata
//
// Example
//
//	engine.GET("/v1/headers", build.Header)
func Header(c echo.Context) error {
	headers := map[string]string{}
	for k, v := range c.Request().Header {
		headers[k] = v[0]
	}

	return c.JSON(http.StatusOK, echo.Map{
		"headers": headers,
		"client": map[string]string{
			"ip": c.RealIP(),
		},
	})
}
