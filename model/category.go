package model

// Category 分类
type Category struct {
	CategoryId   string `json:"category_id" gorm:"column:category_id""`
	CategoryName string `json:"category_name" gorm:"column:category_name"`
	CategoryDesc string `json:"category_desc" gorm:"column:category_desc"`
}
