package handler

import (
	"manageSystem/model"
	"manageSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

// func (h *UserHandler) GetEntity(result model.User) resp.User {
// 	return resp.User{
// 		Id:        result.UserId,
// 		Key:       result.UserId,
// 		UserId:    result.UserId,
// 		NickName:  result.NickName,
// 		Mobile:    result.Mobile,
// 		Address:   result.Address,
// 		IsDeleted: result.IsDeleted,
// 		IsLocked:  result.IsLocked,
// 	}
// }

func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := RespEntity{
		Code:  OperateOk,
		Msg:   OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	u := model.User{
		UserID: userId,
	}
	result, err := h.UserSrv.Get(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = RespEntity{
		Code:  http.StatusOK,
		Msg:   OperateOk.String(),
		Total: 0,
		Data:  *result,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}
