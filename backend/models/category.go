package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Products    []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}
