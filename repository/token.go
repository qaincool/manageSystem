package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"manageSystem/model"
	"manageSystem/model/request"
	"manageSystem/utils"
	"time"
)

type TokenRepository struct {
	DB *gorm.DB
}

type TokenRepoInterface interface {
	AuthLoginUser(loginUser request.LoginReq) (*model.User, error)
	AuthToken(loginToken string) ([]*model.Token, error)
	CreateToken(loginUser request.LoginReq) (*model.Token, error)
	GetTokenUser(loginUser request.LoginReq) (*model.User, error)
}

// AuthLoginUser 验证登录的账号密码是否正确
func (repo *TokenRepository) AuthLoginUser(loginUser request.LoginReq) (*model.User, error) {
	var total int64
	var user = &model.User{
		Mobile:   loginUser.Mobile,
		Password: loginUser.Password,
	}
	repo.DB.First(&user).Count(&total)
	if total > 0 && utils.Md5(loginUser.Password) == user.Password {
		return user, nil
	} else {
		return nil, errors.New("用户不存在或密码错误")
	}
}

func (repo *TokenRepository) AuthToken(loginToken string) ([]*model.Token, error) {
	var total int64
	var userTokens []*model.Token
	repo.DB.Find(&userTokens).Where("token = ?", loginToken).Count(&total)
	if total > 0 {
		return userTokens, nil
	} else {
		return nil, errors.New("token不存在或错误")
	}
}

// CreateToken 为登录用户创建token
func (repo *TokenRepository) CreateToken(loginUser request.LoginReq) (*model.Token, error) {
	var user = &model.User{
		Mobile:   loginUser.Mobile,
		Password: loginUser.Password,
	}
	repo.DB.First(&user)
	var token = &model.Token{
		UserID: user.UserID,
		Token:  utils.CreateToken(user.UserID, time.Now().Add(time.Hour)),
		// token过期时间设置为1小时
		Expiry: time.Now().Add(time.Hour),
	}
	err := repo.DB.Create(&token).Error
	if err != nil {
		return nil, fmt.Errorf("token创建失败")
	}
	return token, err
}

func (repo *TokenRepository) GetTokenUser(loginUser request.LoginReq) (*model.User, error) {
	var user *model.User
	repo.DB.Find(&user).Where("mobile = ?", loginUser.Mobile)
	return user, nil

}
