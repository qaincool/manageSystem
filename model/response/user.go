package response

import "manageSystem/model"

// UserResp 响应体
type UserResp struct {
	UserId   string `json:"id" gorm:"column:user_id"`
	Username string `json:"username" gorm:"column:username"`
	Mobile   string `json:"mobile" gorm:"column:mobile"`
	Address  string `json:"address" gorm:"column:address"`
	RoleId   uint   `json:"role" gorm:"column:role"`
}

func UserModelMapEntity(user *model.User) *UserResp {
	var userEntity = &UserResp{
		UserId:   user.UserID,
		Username: user.Username,
		Mobile:   user.Mobile,
		Address:  user.Address,
		RoleId:   user.RoleId,
	}
	return userEntity
}
