package userHandler

import (
	"com.phpstu/webapp/src/handler"
	"com.phpstu/webapp/src/models/request"
	"com.phpstu/webapp/src/pkg/errno"
	"com.phpstu/webapp/src/service"
	"github.com/gin-gonic/gin"
)

//用户注册
func Register(c *gin.Context) {
	registerForm := request.RegisterForm{}
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		handler.SendJson(c, errno.BindJsonErr, nil)
		return
	}

	serve := service.UserServe{}
	if err := serve.Register(registerForm); err != nil {
		handler.SendJson(c, err, nil)
	} else {
		handler.SendJson(c, errno.OK, nil)
	}

}

//用户登录
func Login(c *gin.Context) {
	loginForm := request.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		handler.SendJson(c, errno.BindJsonErr, nil)
		return
	}
	serve := service.UserServe{}
	if err, token := serve.Login(loginForm); err != nil {
		handler.SendJson(c, err, nil)
	} else {
		handler.SendJson(c, errno.OK, gin.H{
			"token": token,
		})
	}
}
