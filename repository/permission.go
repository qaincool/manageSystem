package repository

import (
	"manageSystem/model"

	"github.com/jinzhu/gorm"
)

type PermissionRepository struct {
	DB *gorm.DB
}

type PermissionRepoInterface interface {
	GetPermissionByRoleId(roleId string) ([]*model.Permission, error)
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
