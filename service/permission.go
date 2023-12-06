package service

import (
	"manageSystem/model"
	"manageSystem/repository"
)

type PermissionSrv interface {
	GetPermissionTree(permissions []*model.Permission) (map[string][]string, error)
}

type PermissionService struct {
	Repo repository.PermissionRepoInterface
}

func (srv PermissionService) GetPermissionTree(roleId string) (map[string][]string, error) {
	return srv.Repo.GetPermissionTree(roleId)
}
