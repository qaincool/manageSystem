package request

import "manageSystem/model"

// UserReq 请求体
type UserReq struct {
	UserId   string `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Address  string `json:"address"`
	RoleId   uint   `json:"role"`
}

func UserModelMapEntity(user *UserReq) *model.User {
	var userEntity = &model.User{
		UserID:   user.UserId,
		Username: user.Username,
		Mobile:   user.Mobile,
		Password: user.Password,
		Address:  user.Address,
		RoleId:   user.RoleId,
	}
	return userEntity
}
