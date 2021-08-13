package io

import (
	"errors"
	"fmt"

	"github.com/hmrkm/simple-auth/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct {
	Conn *gorm.DB
}

func OpenMysql(user string, password string, database string) Mysql {
	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return Mysql{
		Conn: db,
	}
}

func (m Mysql) Close() {
	db, err := m.Conn.DB()
	if err != nil {
		panic(err)
	}

	db.Close()
}

func (m Mysql) Find(dest interface{}, conds string, params ...interface{}) error {
	return m.Conn.Find(dest, conds, params).Error
}

func (m Mysql) First(dest interface{}, conds string, params ...interface{}) error {
	err := m.Conn.First(dest, conds, params).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return usecase.ErrNotFound
	}

	return err
}

func (m Mysql) Create(target interface{}) error {
	return m.Conn.Create(target).Error
}
