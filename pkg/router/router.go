package router

import (
	"github.com/gin-gonic/gin"
)

// SetRoutes 注册路由
func SetRoutes(r *gin.Engine) {
	passportRouter := r.Group("/passport")
	{
		passportRouter.POST("/wx/login", (&passport.HandlerPassport{}).WxLogin)
	}
}
