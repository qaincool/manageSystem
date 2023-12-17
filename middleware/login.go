package middleware

import (
	"github.com/gin-gonic/gin"
	"manageSystem/model/response"
	"manageSystem/service"
	"net/http"
)

var TokenSrv service.TokenService

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		entity := response.RespEntity{
			Code:  response.OperateFail,
			Msg:   response.OperateFail.String(),
			Total: 0,
			Data:  nil,
		}
		token := c.GetHeader("token")
		if err := TokenSrv.AuthToken(token); err != nil {
			entity.Msg = "token验证失败：" + err.Error()
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			c.Abort()
		}
		c.Next()
	}
}
