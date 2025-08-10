package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task4-go-zero/internal/types"
	"task4-go-zero/task4/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化GORM数据库连接
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 自动迁移数据模型
	db.AutoMigrate(&types.User{}, &types.Post{}, &types.Comment{})

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
