package handler

import (
	"manageSystem/model/request"
	"manageSystem/model/response"
	"manageSystem/query"
	"manageSystem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	VideoSrv service.VideoRepoSrv
}

func (h *VideoHandler) VideoListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil {
		entity.Msg = "请求参数错误：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	videoList, err := h.VideoSrv.List(&q)
	total, err := h.VideoSrv.GetTotal()

	if err != nil {
		entity.Msg = "查询视频列表失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: total,
		Data:  videoList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// VideoInfoHandler 获取视频信息
// POST /api/v1/video/getVideo
// data: video_name和video_path
func (h *VideoHandler) VideoInfoHandler(c *gin.Context) {
	var videoInfoReqBody request.VideoReq
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&videoInfoReqBody)
	if err != nil {
		entity.Msg = "请求参数错误：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	videoInfo, err := h.VideoSrv.Get(request.VideoModelMapEntity(&videoInfoReqBody))
	if err != nil {
		entity.Msg = "获取用户失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: 1,
		Data:  videoInfo,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// AddVideoHandler 添加视频信息
// POST /api/v1/video/addVideo
// data: 必填字段 video_name和video_path，可选字段 video_intro,video_detail,video_tag,category_id,create_user
func (h *VideoHandler) AddVideoHandler(c *gin.Context) {
	var videoInfoReqBody request.VideoReq
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&videoInfoReqBody)
	if err != nil {
		entity.Msg = "请求参数错误" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	videoInfo, err := h.VideoSrv.Add(*request.VideoModelMapEntity(&videoInfoReqBody))
	if err != nil {
		entity.Msg = "视频添加失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: 1,
		Data:  videoInfo,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// EditVideoHandler 修改视频信息
// POST /api/v1/video/editVideo
// data: 必传参数：video_id
func (h *VideoHandler) EditVideoHandler(c *gin.Context) {
	var videoInfoReqBody request.VideoReq
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&videoInfoReqBody)
	if err != nil {
		entity.Msg = "请求参数错误" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	videoInfo, err := h.VideoSrv.Edit(*request.VideoModelMapEntity(&videoInfoReqBody))
	if err != nil {
		entity.Msg = "视频修改失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: 1,
		Data:  videoInfo,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

// DeleteVideoHandler 修改视频信息
// POST /api/v1/video/deleteVideo
// data: 必传参数：video_id
func (h *VideoHandler) DeleteVideoHandler(c *gin.Context) {
	var videoInfoReqBody request.VideoReq
	entity := response.RespEntity{
		Code:  response.OperateFail,
		Msg:   response.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindQuery(&videoInfoReqBody)
	if err != nil {
		entity.Msg = "请求参数错误" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	_, err = h.VideoSrv.Delete(*request.VideoModelMapEntity(&videoInfoReqBody))
	if err != nil {
		entity.Msg = "视频删除失败：" + err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	entity = response.RespEntity{
		Code:  http.StatusOK,
		Msg:   "OK",
		Total: 1,
		Data:  "",
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}
