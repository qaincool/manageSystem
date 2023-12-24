package model

// Category 分类
type Category struct {
	VideoId      int    `json:"video_id" gorm:"column:video_id;primaryKey"`
	CategoryId   int    `json:"category_id" gorm:"column:category_id;primaryKey;autoIncrement"`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
}
