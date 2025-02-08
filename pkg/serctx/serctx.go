package serctx

import (
	"gorm.io/gorm"
)

var SerCtx *ServerContext

type ServerContext struct {
	Db *gorm.DB
}

func InitServerContext() {
	// 初始化数据库
	db, _ := initMysql()
	// 初始化全局上下文
	SerCtx = &ServerContext{
		Db: db,
	}
}
