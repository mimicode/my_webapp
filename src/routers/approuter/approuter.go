package approuter

import (
	"com.phpstu/webapp/src/handler/indexHandler"
	"com.phpstu/webapp/src/handler/userHandler"
	"com.phpstu/webapp/src/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAppRouter(g *gin.Engine) {
	nLogin := g.Group("/v1")
	{
		//注册
		nLogin.POST("/user/register", userHandler.Register)
		//登录
		nLogin.POST("/user/login", userHandler.Login)
	}

	login := g.Group("/v1", middlewares.CheckLogin())
	{
		//首页
		login.GET("/index", indexHandler.Index)
	}

}
