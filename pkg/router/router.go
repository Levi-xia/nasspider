package router

import (
	"nasspider/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetRoutes 注册路由
func SetRoutes(r *gin.Engine) {
	pageGroup := r.Group("/")
	{
		pageGroup.GET("/", func(c *gin.Context) {
			r.LoadHTMLFiles("templates/layout/base.html", "templates/index/index.html")
			c.HTML(http.StatusOK, "index.html", nil)
		})
		pageGroup.GET("/login", func(c *gin.Context) {
			r.LoadHTMLFiles("templates/layout/base.html", "templates/login/login.html")
			c.HTML(http.StatusOK, "login.html", nil)
		})
	}

	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/login", app.Login)
		apiGroup.POST("/task/add", app.AddTask)
		apiGroup.POST("/task/edit", app.EditTask)
		apiGroup.POST("/task/trigger", app.TriggerTask)
	}
}
