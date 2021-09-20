package io

import (
	"errors"
	"fmt"

	"github.com/hmrkm/simple-auth/domain"

	"gorm.io/gorm"
)

type Mysql struct {
	Conn GormConn
}

func CreateDSN(user string, password string, database string) (dsn string) {
	return fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
}

func (m Mysql) Close() error {
	db, err := m.Conn.DB()
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func (m Mysql) Find(destAddr interface{}, cond string, params ...interface{}) error {
	return m.Conn.Find(destAddr, cond, params).Error
}

func (m Mysql) First(destAddr interface{}, cond string, params ...interface{}) error {
	err := m.Conn.First(destAddr, cond, params).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}

	return err
}

func (m Mysql) Create(value interface{}) error {
	return m.Conn.Create(value).Error
}

func (m Mysql) IsNotFoundError(err error) bool {
	return errors.Is(gorm.ErrRecordNotFound, err)
}
