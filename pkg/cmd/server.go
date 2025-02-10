package cmd

import (
	"context"
	"fmt"
	"log"
	"nasspider/config"
	"nasspider/pkg/constants"
	"nasspider/pkg/middler"
	"nasspider/pkg/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// 全局中间件
	middler.InitMiddleware(r)
	// 前端项目静态资源
	// 其他静态资源
	//...
	// 注册路由
	router.SetRoutes(r)
	return r
}

func RunServer() {
	r := setupRouter()

	if !config.Conf.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConf(config.Conf.Server.Port, constants.ENV_SERVER_PORT)),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
