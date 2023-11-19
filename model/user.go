package model

import "time"

type User struct {
	UserID    string    `json:"userId" gorm:"column:user_id"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Address   string    `json:"address" gorm:"column:address"`
	IsTeacher bool      `json:"is_teacher" gorm:"column:is_teacher"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at;default:null"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:update_at;default:null"`
}
