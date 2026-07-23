package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role" gorm:"default:admin"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
