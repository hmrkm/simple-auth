package main

import (
	"github.com/hmrkm/simple-auth/adapter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, aa adapter.Auth) {
	e.Use(middleware.CORS())
	g := e.Group("/v1")
	g.POST("/auth", func(c echo.Context) error {
		req := adapter.RequestAuth{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := aa.Auth(req)

		if err != nil {
			return ErrorHandler(c, err)
		}

		return c.JSON(200, res)
	})

	g.POST("/verify", func(c echo.Context) error {
		req := adapter.RequestVerify{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := aa.Verify(req)

		if err != nil {
			return ErrorHandler(c, err)
		}

		return c.JSON(200, res)
	})
}
