package repository

import (
	"manageSystem/model"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	DB *gorm.DB
}

type PermissionRepoInterface interface {
	GetPermissionTree(roleId string) (map[string][]string, error)
}

func (repo *PermissionRepository) GetPermissionByRoleId(roleId string) ([]*model.Permission, error) {
	var permissions []*model.Permission
	db := repo.DB

	var rolePermissions []*model.RolePermission

	if err := db.Find(&rolePermissions, "role_id = ?", roleId).Error; err != nil {
		return nil, err
	}
	for _, rolePermission := range rolePermissions {
		var permission *model.Permission
		if err := db.Find(&permission, "permission_id = ?", rolePermission.PermissionId).Error; err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
func (repo *PermissionRepository) GetPermissionTree(roleId string) (map[string][]string, error) {
	permissions, err := repo.GetPermissionByRoleId(roleId)
	if err != nil {
		return nil, err
	}
	var permissionTreeMap = make(map[string][]string)
	db := repo.DB

	for _, permission := range permissions {
		var childrenNameList []string
		var childrenList []*model.Permission

		if err := db.Find(&childrenList, "parent_id = ?", permission.PermissionId).Error; err != nil {
			return nil, err
		}
		for _, c := range childrenList {
			if c.PermissionName != "" {
				childrenNameList = append(childrenNameList, c.PermissionName)
			}
		}
		permissionTreeMap[permission.PermissionName] = childrenNameList
	}
	return permissionTreeMap, nil
}
