package io

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockTable struct {
	Id   string
	Name string
}

func NewMysqlMock() (Mysql, sqlmock.Sqlmock) {
	db, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: db,
			}), &gorm.Config{})
	mysql := NewMysql(gormDB)
	return mysql, sqlMock
}
