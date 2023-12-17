package request

import "manageSystem/model"

// UserReq 请求体
type UserReq struct {
	UserId   string `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Address  string `json:"address"`
	RoleName string `json:"role_name"`
}

func UserModelMapEntity(user *UserReq) *model.User {
	var userEntity = &model.User{
		UserID:   user.UserId,
		Username: user.Username,
		Mobile:   user.Mobile,
		Password: user.Password,
		Address:  user.Address,
		RoleName: user.RoleName,
	}
	return userEntity
}
