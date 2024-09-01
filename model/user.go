package model

import (
	"gorm.io/gorm"
)

const (
	Admin = iota
	General
	Guest
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"unique"` // 用户名
	Email          string `json:"email" gorm:"unique"`    // 邮箱
	HashedPassword string `json:"-"`                      // 密码
	Salt           string `json:"-"`                      // 盐
	Role           int    `json:"role"`                   // 角色
	Avatar         string `json:"avatar"`                 // 头像 URL
	IsActive       bool   `json:"is_active"`              // 是否激活
}
