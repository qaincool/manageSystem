package middleware

import (
	"github.com/gin-gonic/gin"
	"manageSystem/handler"
	"manageSystem/model"
	"manageSystem/service"
	"net/http"
)

var loginSrv service.LoginService

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginEntity = &model.Login{}
		var entity = &handler.RespEntity{
			Code:  handler.OperateFail,
			Msg:   handler.OperateFail.String(),
			Total: 0,
			Data:  nil,
		}
		err := c.ShouldBindJSON(&loginEntity)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"entity": entity})
			c.Abort()
		}
		if loginEntity.Mobile == "" && loginEntity.Password == "" {
			entity.Msg = "手机号码或密码为空"
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			c.Abort()
		}
		isLogin, err := loginSrv.Auth(loginEntity)
		if err != nil {
			entity.Msg = err.Error()
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			c.Abort()
		}

		if isLogin {
			c.Next()
		} else {
			entity.Msg = "手机号或密码错误"
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			c.Abort()
		}

	}
}
