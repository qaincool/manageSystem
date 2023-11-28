package handler

import (
	"manageSystem/model"
	"manageSystem/query"
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

// UserInfoHandler 获取用户信息
// GET /api/v1/user/getUser
// param: id
func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := RespEntity{
		Code:  OperateFail,
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

// UserListHandler 查询所有用户
// GET /api/v1/user/getUserList
// param: page=1 pageSize=10 可不传
func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := RespEntity{
		Code:  OperateFail,
		Msg:   OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	userList, err := h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}

	entity = RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: total,
		Data:  userList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// AddUserHandler 添加一个用户
// POST /api/v1/user/addUser
// data: 必填字段: mobile,role 非必填: username password address
func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := RespEntity{
		Code:  OperateFail,
		Msg:   OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.UserID == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = OperateOk
	entity.Msg = OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

// EditUserHandler 设置用户
// POST /api/v1/user/editUser
// data: 必填字段user_id 非必填: 需要修改的字段
func (h *UserHandler) EditUserHandler(c *gin.Context) {
	u := model.User{}
	entity := RespEntity{
		Code:  OperateFail,
		Msg:   OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.UserSrv.Edit(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = OperateOk
		entity.Msg = OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}

}

// DeleteUserHandler 删除一个用户
// POST /api/v1/user/deleteUser
// data: 必填字段user_id 其他可以不用填写
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	u := model.User{}
	entity := RespEntity{
		Code:  OperateFail,
		Msg:   OperateFail.String(),
		Total: 0,
		Data:  nil,
	}

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	b, err := h.UserSrv.Delete(u.UserID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = OperateOk
		entity.Msg = OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
