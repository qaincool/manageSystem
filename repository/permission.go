package repository

import (
	"manageSystem/model"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	DB *gorm.DB
}

type PermissionRepoInterface interface {
	GetRoleNameByToken(token string) (*model.User, error)
	GetPermissionByRoleName(roleName string) ([]*model.Permission, error)
}

func (repo *PermissionRepository) GetRoleNameByToken(userToken string) (*model.User, error) {
	var user *model.User
	db := repo.DB
	var token = &model.Token{
		Token: userToken,
	}
	db.First(token)
	if err := db.Find(&user, "user_id = ?", token.UserID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *PermissionRepository) GetPermissionByRoleName(roleName string) ([]*model.Permission, error) {
	var permissions []*model.Permission
	db := repo.DB

	var rolePermissions []*model.RolePermission

	var role model.Role
	if err := db.Find(&role, "role_name = ?", roleName).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&rolePermissions, "role_id = ?", role.RoleId).Error; err != nil {
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
