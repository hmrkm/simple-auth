package main

import (
	"time"

	"github.com/hmrkm/simple-auth/adapter"
	"github.com/hmrkm/simple-auth/domain"
	"github.com/hmrkm/simple-auth/io"
	"github.com/hmrkm/simple-auth/usecase"
	"github.com/kelseyhightower/envconfig"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func main() {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysqlDriver.Open(io.CreateDSN(
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlDatabase,
	)))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(config.MysqlMaxIdelConns)
	sqlDB.SetMaxOpenConns(config.MysqlMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	mysql := io.NewMysql(db)
	defer mysql.Close()

	usd := domain.NewUserService(mysql)
	tsd := domain.NewTokenService(mysql)
	au := usecase.NewAuth(usd, tsd)
	tu := usecase.NewToken(mysql)
	aa := adapter.NewAuth(au, tu, config.TokenExpireHour)

	e := echo.New()
	Route(e, aa)

	e.Logger.Fatal(e.Start(":80"))
}
