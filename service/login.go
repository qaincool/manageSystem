package service

import (
	"manageSystem/repository"
)

type LoginSrv interface {
	Auth() (map[string]string, error)
}

type LoginService struct {
	Repo repository.LoginRepository
}

func (srv *LoginService) Auth() (map[string]string, error) {
	return srv.Repo.Auth()
}
