package service

import (
	"manageSystem/model"
	"manageSystem/repository"
)

type LoginSrv interface {
	Auth(login *model.Login) (bool, error)
}

type LoginService struct {
	Repo repository.LoginRepository
}

func (srv *LoginService) Auth(login *model.Login) (bool, error) {
	return srv.Repo.Auth(login)
}
