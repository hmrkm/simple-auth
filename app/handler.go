package main

import (
	"errors"

	"github.com/hmrkm/simple-auth/domain"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err error) (json error) {

	if errors.As(err, &domain.ErrNotFound) ||
		errors.As(err, &domain.ErrTokenWasExpired) {
		return c.JSON(404, nil)
	}

	if errors.As(err, &domain.ErrInvalidVerify) {
		return c.JSON(401, nil)
	}

	if err != nil {
		return c.JSON(500, nil)
	}

	return nil
}
