package model

type Role struct {
	RoleId   string `json: role_id gorm:"column:role_id"`
	RoleName string `json: role_name gorm:"role_name"`
	RoleDesc string `json: role_desc gorm:"role_desc"`
}

type Permission struct {
	PermissionId   string `json: permission_id gorm:"permission_id"`
	PermissionName string `json: permission_name gorm:"permission_name"`
	PermissionDesc string `json: permission_desc gorm:"permission_desc"`
}

type RolePermission struct {
	RoleId       string `json: role_id gorm:"column:role_id"`
	PermissionId string `json: permission_id gorm:"permission_id"`
}
