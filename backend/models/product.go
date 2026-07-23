package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name" gorm:"not null"`
	Slug        string   `json:"slug" gorm:"uniqueIndex;not null"`
	Description string   `json:"description"`
	Price       float64  `json:"price" gorm:"not null"`
	SalePrice   *float64 `json:"sale_price"`
	Images      []string `json:"images" gorm:"serializer:json"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Sizes       []string `json:"sizes" gorm:"serializer:json"`
	Colors      []string `json:"colors" gorm:"serializer:json"`
	Stock       int      `json:"stock" gorm:"default:100"`
	Featured    bool     `json:"featured" gorm:"default:false"`
	Tags        []string `json:"tags" gorm:"serializer:json"`
}
