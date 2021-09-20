package io

import (
	"errors"
	"fmt"

	"github.com/hmrkm/simple-auth/domain"

	"gorm.io/gorm"
)

type Mysql struct {
	conn GormConn
}

func NewMysql(conn GormConn) Mysql {
	return Mysql{
		conn: conn,
	}
}

func CreateDSN(user string, password string, database string) (dsn string) {
	return fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)
}

func (m Mysql) Close() error {
	db, err := m.conn.DB()
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func (m Mysql) Find(destAddr interface{}, cond string, params ...interface{}) error {
	return m.conn.Find(destAddr, cond, params).Error
}

func (m Mysql) First(destAddr interface{}, cond string, params ...interface{}) error {
	err := m.conn.First(destAddr, cond, params).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}

	return err
}

func (m Mysql) Create(value interface{}) error {
	return m.conn.Create(value).Error
}

func (m Mysql) IsNotFoundError(err error) bool {
	return errors.Is(gorm.ErrRecordNotFound, err)
}
