package request

import "manageSystem/model"

// VideoReq 请求体
type VideoReq struct {
	VideoId   string `json:"video_id"`
	VideoName string `json:"video_name"`
	// 简介
	VideoIntro string `json:"video_intro"`
	// 视频存放路径
	VideoPath string `json:"video_path"`
	// 详细描述
	VideoDetail string `json:"video_detail"`
	// 标签信息
	VideoTag []string `json:"video_tag"`
	// 所属品类id号
	CategoryId string `json:"category_id"`
	// 上传用户
	CreateUser string `json:"create_user"`
}

func VideoModelMapEntity(video *VideoReq) *model.Video {
	var videoEntity = &model.Video{
		VideoId:     video.VideoId,
		VideoName:   video.VideoName,
		VideoIntro:  video.VideoIntro,
		VideoPath:   video.VideoPath,
		VideoDetail: video.VideoDetail,
		VideoTag:    video.VideoTag,
		CategoryId:  video.CategoryId,
		CreateUser:  video.CreateUser,
	}
	return videoEntity
}
