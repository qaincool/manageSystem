package model

import "time"

// User 用户
type User struct {
	UserID   string    `json:"userId" gorm:"column:user_id;primary_key;not null"`
	Username string    `json:"username" gorm:"column:username"`
	Password string    `json:"password" gorm:"column:password"`
	Mobile   string    `json:"mobile" gorm:"column:mobile"`
	Address  string    `json:"address" gorm:"column:address"`
	RoleName string    `json:"role_name" gorm:"column:role_name"`
	Email    string    `json:"email" gorm:"column:email"`
	QQ       string    `json:"qq"  gorm:"column:qq"`
	Age      uint      `json:"age" gorm:"column:age"`
	Sex      string    `json:"sex" gorm:"column:sex"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at;default:null"`
}
