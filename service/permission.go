package service

import (
	"manageSystem/model"
	"manageSystem/model/response"
	"manageSystem/repository"
)

type PermissionSrv interface {
	GetRoleNameByToken(loginUser string) (*model.User, error)
	GetPermissionByRoleName(roleId string) ([]*response.UserPermission, error)
}

type PermissionService struct {
	Repo repository.PermissionRepoInterface
}

func (srv PermissionService) GetRoleNameByToken(loginUser string) (*model.User, error) {
	return srv.Repo.GetRoleNameByToken(loginUser)
}

func (srv PermissionService) GetPermissionByRoleName(roleId string) ([]*response.UserPermission, error) {
	permissionsEntities, err := srv.Repo.GetPermissionByRoleName(roleId)
	if err != nil {
		return nil, err
	}

	var userPermissions []*response.UserPermission
	for _, permissionEntity := range permissionsEntities {
		var userPermission = &response.UserPermission{
			Name:        permissionEntity.PermissionName,
			Description: permissionEntity.PermissionDesc,
			Level:       permissionEntity.PermissionLevel,
		}
		userPermissions = append(userPermissions, userPermission)
	}
	return userPermissions, nil
}
