package repository

import (
	"errors"
	"gorm.io/gorm"
	"manageSystem/model"
	"manageSystem/utils"
)

type LoginRepository struct {
	DB *gorm.DB
}

type LoginRepoInterface interface {
	Auth(login *model.Login) (bool, error)
}

func (repo *LoginRepository) Auth(login *model.Login) (bool, error) {
	var user = &model.User{}
	repo.DB.Find(&user).Where("mobile = ?", login.Mobile)
	if user.Password == "" {
		return false, errors.New("用户不存在")
	}
	if user.Password != utils.Md5(login.Password) {
		return false, errors.New("用户名或密码错误")
	}
	return true, nil

}
