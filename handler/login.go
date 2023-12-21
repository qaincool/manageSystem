package handler

import (
	"github.com/gin-gonic/gin"
	"manageSystem/model/request"
	"manageSystem/model/response"
	"manageSystem/service"
	"net/http"
)

type LoginHandler struct {
	TokenSrv service.TokenSrv
	UserSrv  service.UserSrv
}

func (h *LoginHandler) LoginHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}

	login := request.LoginReq{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		entity.Msg = "请求体格式错误"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	token, err := h.TokenSrv.CreateToken(login)
	if err != nil {
		entity.Msg = "token创建失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	user, err := h.TokenSrv.GetTokenUser(login)
	userToken := response.LoginModelMapEntity(token)
	userToken.User = *response.UserModelMapEntity(user)

	entity.Code = response.OperateOk
	entity.Msg = "token创建成功"
	entity.Data = userToken
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}
