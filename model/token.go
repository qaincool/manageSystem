package model

import "time"

type Token struct {
	Token  string    `json:"token" gorm:"column:token;primary_key;not null"`
	UserID string    `json:"userId" gorm:"column:user_id"`
	Expiry time.Time `json:"expiry" gorm:"column:expiry"`
}
