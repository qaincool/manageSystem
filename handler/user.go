package handler

import (
	"manageSystem/model/request"
	"manageSystem/model/response"
	"manageSystem/query"
	"manageSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

// UserInfoHandler 获取用户信息
// GET /api/v1/user/getUser
// param: id
func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := request.UserReq{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	userInfo, err := h.UserSrv.Get(*request.UserModelMapEntity(&u))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   response.OperateOk.String(),
		Total: 0,
		Data:  response.UserModelMapEntity(userInfo),
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// UserListHandler 查询所有用户
// GET /api/v1/user/getUserList
// param: page=1 pageSize=10 可不传
func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	userInfoList, err := h.UserSrv.List(&q)

	if err != nil {
		panic(err)
	}

	var userEntityList []*response.UserResp
	for _, userInfo := range userInfoList {
		userEntityList = append(userEntityList, response.UserModelMapEntity(userInfo))
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: int64(len(userEntityList)),
		Data:  userEntityList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// AddUserHandler 添加一个用户
// POST /api/v1/user/addUser
// data: 必填字段: mobile,role 非必填: username password address
func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := request.UserReq{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		entity.Msg = "请求体格式错误"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.Add(request.UserModelMapEntity(&u))
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if r.UserID == "" {
		entity.Msg = "用户插入失败"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity.Code = response.OperateOk
	entity.Msg = "用户注册成功"
	entity.Data = r
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

// EditUserHandler 设置用户
// POST /api/v1/user/editUser
// data: 必填字段user_id 非必填: 需要修改的字段
func (h *UserHandler) EditUserHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := request.UserReq{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		entity.Msg = "请求体格式错误"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.UserSrv.Edit(*request.UserModelMapEntity(&u))
	if err != nil {
		entity.Msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = response.OperateOk
		entity.Msg = "用户更新成功"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}

}

// DeleteUserHandler 删除一个用户
// POST /api/v1/user/deleteUser
// data: 必填字段user_id 其他可以不用填写
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := request.UserReq{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	b, err := h.UserSrv.Delete(u.UserId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = response.OperateOk
		entity.Msg = "用户删除成功"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
