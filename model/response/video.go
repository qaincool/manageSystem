package response

import (
	"manageSystem/model"
	"manageSystem/utils"
	"time"
)

// VideoResp 请求体
type VideoResp struct {
	VideoId   int    `json:"id"`
	VideoName string `json:"name"`
	// 简介
	VideoIntro string `json:"intro"`
	// 视频存放路径
	VideoPath string `json:"path"`
	// 详细描述
	VideoDetail string `json:"detail"`
	// 标签信息
	VideoTag []string `json:"tag"`
	// 所属品类id号
	Category []string `json:"categories"`
	// 上传用户
	CreateUser string `json:"create_user"`
	// 时间上传时间
	CreateTime time.Time `json:"create_time"`
}

func VideoModelMapEntity(video *model.Video) *VideoResp {
	var videoEntity = &VideoResp{
		VideoId:     video.VideoId,
		VideoName:   video.VideoName,
		VideoIntro:  video.VideoIntro,
		VideoPath:   video.VideoPath,
		VideoDetail: video.VideoDetail,
		VideoTag:    utils.StringToArray(video.VideoTag),
		Category:    CategoryModelMapEntity(video.Category),
		CreateUser:  video.CreateUser,
		CreateTime:  video.CreateTime,
	}
	return videoEntity
}

func CategoryModelMapEntity(CategoryModels []*model.Category) []string {
	var categories []string
	for _, categoryModel := range CategoryModels {
		categories = append(categories, categoryModel.CategoryName)
	}
	return categories
}
