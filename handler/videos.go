package handler

import (
	"github.com/gin-gonic/gin"
	"manageSystem/model"
	"manageSystem/model/request"
	"manageSystem/model/response"
	"manageSystem/query"
	"manageSystem/service"
	"net/http"
	"time"
)

type VideoHandler struct {
	VideoSrv service.VideoService
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
	var videoInfoReqBody model.Video
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

	videoInfoReqBody.CreateTime = time.Now()
	videoInfo, err := h.VideoSrv.Add(videoInfoReqBody)
	if err != nil {
		entity.Msg = "用户添加失败：" + err.Error()
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
