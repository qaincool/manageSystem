package main

import (
	"manageSystem/utils"
	"testing"
	"time"
)

// Category 分类
type Category struct {
	VideoId      uint   `json:"video_id" gorm:"column:video_id"`
	CategoryId   uint   `json:"category_id" gorm:"column:category_id;primaryKey;autoIncrement"`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
	CategoryDesc string `json:"category_desc" gorm:"column:category_desc"`
}

// Video 视频类
type Video struct {
	VideoId   uint   `json:"video_id" gorm:"column:video_id;primary_key;not null;autoIncrement"`
	VideoName string `json:"video_name" gorm:"column:video_name"`
	// 简介
	VideoIntro string `json:"video_intro" gorm:"column:video_intro"`
	// 视频存放路径
	VideoPath string `json:"video_path" gorm:"column:video_path"`
	// 详细描述
	VideoDetail string `json:"video_detail" gorm:"column:video_detail"`
	// 标签信息
	VideoTag string `json:"video_tag" column:"video_tag"`
	// 上传用户
	CreateUser string `json:"create_user" gorm:"column:create_user"`
	// 上传时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`

	// 添加一对多的关联关系
	Category []Category `json:"category" gorm:"foreignKey:video_id"`
}

func TestMysql(t *testing.T) {
	// 创建数据库连接
	utils.InitDB()
	db := utils.DB
	db.AutoMigrate(&Video{}, &Category{})

}
