package service

import (
	"errors"
	"manageSystem/model"
	"manageSystem/repository"
)

type PermissionSrv interface {
	GetPermissionByRoleId(roleId string) ([]*model.Permission, error)
}

type PermissionService struct {
	Repo repository.PermissionRepoInterface
}

func (srv PermissionService) GetPermissionByRoleId(roleId string) ([]*model.Permission, error) {
	if roleId == "" {
		return nil, errors.New("role is null")
	}
	return srv.Repo.GetPermissionByRoleId(roleId)
}
