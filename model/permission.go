package model

type Role struct {
	RoleId   uint   `json:"role_id" gorm:"column:role_id;primary_key;not null"`
	RoleName string `json:"role_name" gorm:"column:role_name"`
	RoleDesc string `json:"role_desc" gorm:"column:role_desc"`
}

type Permission struct {
	PermissionId    uint   `json:"permission_id" gorm:"column:permission_id;primary_key;not null"`
	PermissionName  string `json:"permission_name" gorm:"column:permission_name"`
	PermissionDesc  string `json:"permission_desc" gorm:"column:permission_desc"`
	PermissionLevel string `json:"permission_level" gorm:"column:permission_level"`
	ParentId        uint   `json:"parent_id" gorm:"column:parent_id"`
}

type RolePermission struct {
	Id           uint `json:"id" gorm:"column:id;primary_key;not null"`
	RoleId       uint `json:"role_id" gorm:"column:role_id"`
	PermissionId uint `json:"permission_id" gorm:"permission_id"`
}
