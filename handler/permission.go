package handler

import (
	"github.com/gin-gonic/gin"
	"manageSystem/model/response"
	"manageSystem/service"
	"net/http"
)

type PermissionHandler struct {
	PermissionSrv service.PermissionSrv
}

// PermissionHandler 根据登录的用户名获取用户权限
// GET /api/v1/user/permission
func (h *PermissionHandler) PermissionHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	userToken := c.GetHeader("token")
	user, err := h.PermissionSrv.GetRoleNameByToken(userToken)
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	userPermissions, err := h.PermissionSrv.GetPermissionByRoleName(user.RoleName)
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   response.OperateOk.String(),
		Total: 0,
		Data:  response.PermissionModelMapEntity(user.Username, user.RoleName, userPermissions),
	}

	c.JSON(http.StatusOK, gin.H{"entity": entity})
}
