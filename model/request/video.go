package request

import (
	"manageSystem/model"
	"manageSystem/utils"
)

// VideoReq 请求体
type VideoReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// 简介
	Intro string `json:"intro"`
	// 视频存放路径
	Path string `json:"path"`
	// 详细描述
	Detail string `json:"detail"`
	// 标签信息
	Tag      []string `json:"tag"`
	Category []string `json:"categories"`
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
		Category:    CategoryModelMapEntity(video.Id, video.Category),
		CreateUser:  video.CreateUser,
	}
	return videoEntity
}

func CategoryModelMapEntity(VideoId int, categories []string) []*model.Category {
	var CategoryModels []*model.Category
	for _, category := range categories {
		CategoryModels = append(CategoryModels, &model.Category{
			VideoId:      VideoId,
			CategoryName: category,
		})
	}
	return CategoryModels
}
