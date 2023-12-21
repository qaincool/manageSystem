package model

import "time"

// Video 视频类
type Video struct {
	VideoId   string `json:"video_id" gorm:"column:video_id;primary_key;not null"`
	VideoName string `json:"video_name" gorm:"column:video_name"`
	// 简介
	VideoIntro string `json:"video_intro" gorm:"column:video_intro"`
	// 视频存放路径
	VideoPath string `json:"video_path" gorm:"column:video_path"`
	// 详细描述
	VideoDetail string `json:"video_detail" gorm:"column:video_detail"`
	// 标签信息
	VideoTag string `json:"video_tag" column:"video_tag"`
	// 所属品类id号
	CategoryId string `json:"category_id" gorm:"column:category_id"`
	// 上传用户
	CreateUser string `json:"create_user" gorm:"column:create_user"`
	// 上传时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
