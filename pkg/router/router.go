package router

import (
	"nasspider/pkg/app"
	"nasspider/pkg/middler"

	"github.com/gin-gonic/gin"
)

// SetRoutes 注册路由
func SetRoutes(r *gin.Engine) {
	// 静态文件
	r.Static("/static", "static/")

	pageGroup := r.Group("/")
	{
		pageGroup.GET("/", middler.NeedLoginHandler(), func(c *gin.Context) {
			r.LoadHTMLFiles("templates/layout/base.html", "templates/index/index.html")
			app.Index(c)
		})
		pageGroup.GET("/login", func(c *gin.Context) {
			r.LoadHTMLFiles("templates/layout/base.html", "templates/login/login.html")
			app.Login(c)
		})
	}
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/login/submit", app.SubmitLogin)
		apiGroup.POST("/task/list", middler.NeedLoginHandler(), app.GetTaskList)
		apiGroup.POST("/task/add", middler.NeedLoginHandler(), app.AddTask)
		apiGroup.POST("/task/edit", middler.NeedLoginHandler(), app.EditTask)
		apiGroup.POST("/task/trigger", middler.NeedLoginHandler(), app.TriggerTask)
		apiGroup.POST("/task/triggerAll", middler.NeedLoginHandler(), app.TriggerAllTask)
		apiGroup.POST("/download/add", middler.NeedLoginHandler(), app.AddDownload)
	}
}
