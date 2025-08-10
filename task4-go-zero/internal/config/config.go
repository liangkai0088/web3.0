package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Log   LogConf
	Mysql MysqlConf
	Auth  AuthConf
}

type LogConf struct {
	Path       string
	Level      string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type MysqlConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Timeout  int
}

type AuthConf struct {
	JwtSecret   string
	TokenExpiry int
}
