package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`

	RoleID uint
}

type Role struct {
	gorm.Model
	Name        string       `json:"name"`
	User        []User       `json:"users"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}
