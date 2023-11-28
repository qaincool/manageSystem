package middleware

import (
	"github.com/gin-gonic/gin"
	"manageSystem/handler"
	"manageSystem/model"
	"net/http"
)

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
			return
		}
		if loginEntity.Mobile == "" && loginEntity.Password == "" {
			entity.Msg = "手机号码或密码为空"
			c.JSON(http.StatusUnauthorized, gin.H{"entity": entity})
			return
		}

	}
}
