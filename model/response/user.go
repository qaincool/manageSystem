package response

import "manageSystem/model"

// UserResp 响应体
type UserResp struct {
	UserId   string `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	RoleName string `json:"role_name"`
	Email    string `json:"email"`
	QQ       string `json:"qq"`
	Age      uint   `json:"age"`
	Sex      string `json:"sex"`
}

func UserModelMapEntity(user *model.User) *UserResp {
	var userEntity = &UserResp{
		UserId:   user.UserID,
		Username: user.Username,
		Mobile:   user.Mobile,
		Address:  user.Address,
		RoleName: user.RoleName,
		Email:    user.Email,
		QQ:       user.QQ,
		Age:      user.Age,
		Sex:      user.Sex,
	}
	return userEntity
}
