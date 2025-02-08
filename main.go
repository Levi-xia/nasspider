package main

import (
	"nasspider/config"
	"nasspider/pkg/cmd"
	"nasspider/pkg/common"
	"nasspider/pkg/cron"
	"nasspider/pkg/logger"
	"nasspider/pkg/serctx"
)

func main() {

	// 初始化配置
	config.InitConfig()
	// 初始化日志
	logger.InitLog()
	// 初始化服务
	serctx.InitServerContext()
	// 初始化验证器
	common.InitValidator()
	// 初始化定时任务
	cron.InitCron()
	// 启动服务
	cmd.RunServer()
}
