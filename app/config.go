package main

type Config struct {
	MysqlUser       string `split_words:"true"`
	MysqlPassword   string `split_words:"true"`
	MysqlDatabase   string `split_words:"true"`
	MysqlMaxIdelConns int `split_words:"true"`
	MysqlMaxOpenConns int `split_words:"true"`
	TokenExpireHour int    `split_words:"true"`
}
