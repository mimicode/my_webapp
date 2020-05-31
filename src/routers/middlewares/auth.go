package middlewares

import (
	"com.phpstu/webapp/src/handler"
	"com.phpstu/webapp/src/pkg/errno"
	"com.phpstu/webapp/src/pkg/jwt"
	"com.phpstu/webapp/src/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		signContext, err := jwt.ParseRequest(c, viper.GetString("jwt"))
		if err != nil {
			handler.SendJson(c, errno.RequiredLoginError, nil)
			return
		}

		serve := service.UserServe{}
		userInfo := serve.GetUserInfoById(int(signContext.ID))
		if userInfo == nil {
			handler.SendJson(c, errno.RequiredLoginError, nil)
			return
		}

		c.Set("userInfo", userInfo)
		c.Next()
	}
}
