package service

import (
	"errors"
	"manageSystem/model"
	"manageSystem/model/request"
	"manageSystem/repository"
	"time"
)

type TokenSrv interface {
	AuthToken(loginToken string) error
	CreateToken(loginUser request.LoginReq) (*model.Token, error)
}

type TokenService struct {
	Repo repository.TokenRepoInterface
}

func (srv *TokenService) AuthToken(loginToken string) error {
	userTokens, err := srv.Repo.AuthToken(loginToken)
	if err != nil {
		return err
	}
	for _, token := range userTokens {
		// 如果过期时间比现在的时间晚，则说明token有效直接返回
		if token.Expiry.After(time.Now()) {
			return nil
		}
	}
	return errors.New("token已经过期")
}

func (srv *TokenService) CreateToken(loginUser request.LoginReq) (*model.Token, error) {
	return srv.Repo.CreateToken(loginUser)
}
