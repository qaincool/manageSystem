package model

type Role struct {
	RoleId   uint   `json:"role_id" gorm:"column:role_id"`
	RoleName string `json:"role_name" gorm:"column:role_name"`
	RoleDesc string `json:"role_desc" gorm:"column:role_desc"`
}

type Permission struct {
	PermissionId   uint   `json:"permission_id" gorm:"column:permission_id"`
	PermissionName string `json:"permission_name" gorm:"column:permission_name"`
	PermissionDesc string `json:"permission_desc" gorm:"column:permission_desc"`
	ParentId       uint   `json:"parent_id" gorm:"column:parent_id"`
}

type RolePermission struct {
	RoleId       uint `json:"role_id" gorm:"column:role_id"`
	PermissionId uint `json:"permission_id" gorm:"permission_id"`
}
