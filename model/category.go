package model

type Category struct {
	CategoryId string `json:"category_id" gorm:"column:category_id""`
	Name       string `json:"name" gorm:"column:name"`
}
