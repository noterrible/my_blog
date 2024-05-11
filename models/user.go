package models

import "gorm.io/gorm"

type Role int

const (
	PermissionUser  Role = 0
	PermissionBlack Role = 1
	PermissionAdmin Role = 2
)

type User struct {
	gorm.Model
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Role     Role   `json:"role"`
	Token    string `json:"token"` //其他平台的token

}

func ParseRole(role Role) string {
	switch role {
	case PermissionUser:
		return "普通用户"
	case PermissionBlack:
		return "黑名单"
	case PermissionAdmin:
		return "管理员"
	default:
		return "未知角色"
	}

}
