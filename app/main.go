package main

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/hmrkm/simple-auth/adapter"
	"github.com/hmrkm/simple-auth/io"
	"github.com/hmrkm/simple-auth/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	tokenExpireHour, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_HOUR"))
	if err != nil {
		panic(err)
	}

	mysql := io.OpenMysql(mysqlUser, mysqlPassword, mysqlDatabase)
	defer mysql.Close()

	usu := usecase.NewUserService(mysql)
	tsu := usecase.NewTokenService(mysql)
	aa := adapter.NewAuthAdapter(usu, tsu)
	ta := adapter.NewTokenAdapter(mysql)

	e := echo.New()
	e.POST("/v1/auth", func(c echo.Context) error {
		req := adapter.RequestAuth{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := aa.Verify(req, time.Now(), tokenExpireHour)

		if errors.Is(usecase.ErrNotFound, err) {
			return c.JSON(404, nil)
		}

		if errors.Is(usecase.ErrInvalidVerify, err) {
			return c.JSON(401, nil)
		}

		if err != nil {
			return c.JSON(500, nil)
		}

		return c.JSON(200, res)
	})

	e.POST("/v1/verify", func(c echo.Context) error {
		req := adapter.RequestVerify{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := ta.Verify(req, time.Now())

		if errors.Is(err, usecase.ErrNotFound) ||
			errors.Is(err, usecase.ErrTokenWasExpired) {
			return c.JSON(404, nil)
		}

		if err != nil {
			return c.JSON(500, nil)
		}

		return c.JSON(200, res)
	})

	e.Logger.Fatal(e.Start(":80"))
}
