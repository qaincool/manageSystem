package request

import (
	"manageSystem/model"
	"manageSystem/utils"
)

// VideoReq 请求体
type VideoReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// 简介
	Intro string `json:"intro"`
	// 视频存放路径
	Path string `json:"path"`
	// 详细描述
	Detail string `json:"detail"`
	// 标签信息
	Tag []string `json:"tag"`
	// 所属品类id号
	CategoryId string `json:"category_id"`
	// 上传用户
	CreateUser string `json:"create_user"`
}

func VideoModelMapEntity(video *VideoReq) *model.Video {
	var videoEntity = &model.Video{
		VideoId:     video.Id,
		VideoName:   video.Name,
		VideoIntro:  video.Intro,
		VideoPath:   video.Path,
		VideoDetail: video.Detail,
		VideoTag:    utils.ArrayToString(video.Tag),
		CategoryId:  video.CategoryId,
		CreateUser:  video.CreateUser,
	}
	return videoEntity
}
