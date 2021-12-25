package main

import (
	"os"
	"strconv"

	"github.com/hmrkm/simple-auth/adapter"
	"github.com/hmrkm/simple-auth/domain"
	"github.com/hmrkm/simple-auth/io"
	"github.com/hmrkm/simple-auth/usecase"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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

	db, err := gorm.Open(mysqlDriver.Open(io.CreateDSN(mysqlUser, mysqlPassword, mysqlDatabase)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	mysql := io.NewMysql(db)
	defer mysql.Close()

	usd := domain.NewUserService(mysql)
	tsd := domain.NewTokenService(mysql)
	au := usecase.NewAuth(usd, tsd)
	tu := usecase.NewToken(mysql)
	aa := adapter.NewAuth(au, tu, tokenExpireHour)

	e := echo.New()
	Route(e, aa)

	e.Logger.Fatal(e.Start(":80"))
}
