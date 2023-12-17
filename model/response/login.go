package response

import (
	"manageSystem/model"
	"time"
)

// LoginResp 响应体
type LoginResp struct {
	Token  string    `json:"token"`
	UserID string    `json:"userId"`
	Expiry time.Time `json:"expiry"`
}

func LoginModelMapEntity(token *model.Token) *LoginResp {
	var loginEntity = &LoginResp{
		UserID: token.UserID,
		Token:  token.Token,
		Expiry: token.Expiry,
	}
	return loginEntity
}
