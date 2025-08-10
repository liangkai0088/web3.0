package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql MysqlConfig
	Auth  AuthConfig
}

type MysqlConfig struct {
	DataSource string
}

type AuthConfig struct {
	JwtSecret   string
	TokenExpiry int
}
