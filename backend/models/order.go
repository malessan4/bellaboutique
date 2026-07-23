package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName    string      `json:"customer_name"`
	CustomerEmail   string      `json:"customer_email"`
	CustomerPhone   string      `json:"customer_phone"`
	ShippingAddress string      `json:"shipping_address"`
	City            string      `json:"city"`
	Province        string      `json:"province"`
	PostalCode      string      `json:"postal_code"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Subtotal        float64     `json:"subtotal"`
	ShippingCost    float64     `json:"shipping_cost"`
	Total           float64     `json:"total"`
	Status          string      `json:"status" gorm:"default:pending"`
	MPPaymentID     string      `json:"mp_payment_id"`
	MPPreferenceID  string      `json:"mp_preference_id"`
	Notes           string      `json:"notes"`
}

type OrderItem struct {
	gorm.Model
	OrderID     uint    `json:"order_id"`
	ProductID   uint    `json:"product_id"`
	Product     Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}
