package response

// PermissionResp 用户权限响应体
type PermissionResp struct {
	User        string            `json:"user"`
	RoleName    string            `json:"role_name"`
	Permissions []*UserPermission `json:"permissions"`
}

// UserPermission 请求返回的用户权限字段
type UserPermission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

func PermissionModelMapEntity(loginUser string, roleName string, userPermissions []*UserPermission) *PermissionResp {
	var permissionEntity = &PermissionResp{
		User:        loginUser,
		RoleName:    roleName,
		Permissions: userPermissions,
	}

	return permissionEntity
}
