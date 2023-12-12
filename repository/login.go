package repository

import (
	"gorm.io/gorm"
	"manageSystem/model"
)

type LoginRepository struct {
	DB *gorm.DB
}

type LoginRepoInterface interface {
	Auth() (map[string]string, error)
}

func (repo *LoginRepository) Auth() (map[string]string, error) {
	var userPassMap = make(map[string]string)
	var users []model.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	// 将数据库中存储的账号名/密码存入内存
	for _, user := range users {
		userPassMap[user.Mobile] = user.Password
	}
	return userPassMap, nil
}
