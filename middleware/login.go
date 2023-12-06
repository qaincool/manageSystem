package middleware

import (
	"github.com/gin-gonic/gin"
	"manageSystem/model"
	"manageSystem/resp"
	"manageSystem/service"
	"net/http"
)

var LoginSrv service.LoginService

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginEntity = &model.Login{}
		var entity = &resp.RespEntity{
			Code:  resp.OperateFail,
			Msg:   resp.OperateFail.String(),
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
		isLogin, err := LoginSrv.Auth(loginEntity)

		if isLogin {
			c.Next()
		} else {
			entity.Msg = "手机号或密码错误"
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			c.Abort()
		}

		c.MustGet(gin.AuthUserKey).(string)

	}
}
