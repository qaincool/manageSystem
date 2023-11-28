package model

import "time"

// User 用户
type User struct {
	UserID   string    `json:"userId" gorm:"column:user_id"`
	Username string    `json:"username" gorm:"column:username"`
	Password string    `json:"password" gorm:"column:password"`
	Mobile   string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Address  string    `json:"address" gorm:"column:address"`
	RoleId   uint      `json:"role" gorm:"column:role"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at;default:null"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at;default:null"`
}
