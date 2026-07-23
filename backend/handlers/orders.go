package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"bellaboutique/config"
	"bellaboutique/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewOrderHandler(db *gorm.DB, cfg *config.Config) *OrderHandler {
	return &OrderHandler{db: db, cfg: cfg}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req struct {
		CustomerName    string `json:"customer_name" binding:"required"`
		CustomerEmail   string `json:"customer_email" binding:"required,email"`
		CustomerPhone   string `json:"customer_phone"`
		ShippingAddress string `json:"shipping_address" binding:"required"`
		City            string `json:"city" binding:"required"`
		Province        string `json:"province" binding:"required"`
		PostalCode      string `json:"postal_code"`
		Notes           string `json:"notes"`
		Items []struct {
			ProductID uint   `json:"product_id" binding:"required"`
			Size      string `json:"size"`
			Color     string `json:"color"`
			Quantity  int    `json:"quantity" binding:"required,min=1"`
		} `json:"items" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var orderItems []models.OrderItem
	var subtotal float64
	for _, ir := range req.Items {
		var product models.Product
		if err := h.db.First(&product, ir.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Producto %d no encontrado", ir.ProductID)})
			return
		}
		price := product.Price
		if product.SalePrice != nil {
			price = *product.SalePrice
		}
		sub := price * float64(ir.Quantity)
		subtotal += sub
		orderItems = append(orderItems, models.OrderItem{
			ProductID:   ir.ProductID,
			ProductName: product.Name,
			Size:        ir.Size,
			Color:       ir.Color,
			Price:       price,
			Quantity:    ir.Quantity,
			Subtotal:    sub,
		})
	}

	shipping := 2500.0
	if subtotal >= 30000 {
		shipping = 0
	}

	order := models.Order{
		CustomerName:    req.CustomerName,
		CustomerEmail:   req.CustomerEmail,
		CustomerPhone:   req.CustomerPhone,
		ShippingAddress: req.ShippingAddress,
		City:            req.City,
		Province:        req.Province,
		PostalCode:      req.PostalCode,
		Subtotal:        subtotal,
		ShippingCost:    shipping,
		Total:           subtotal + shipping,
		Notes:           req.Notes,
		Status:          "pending",
	}
	if err := h.db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la orden"})
		return
	}
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}
	h.db.CreateInBatches(&orderItems, 20)
	order.Items = orderItems
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if h.db.Preload("Items").First(&order, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	var orders []models.Order
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit
	var total int64
	h.db.Model(&models.Order{}).Count(&total)
	h.db.Preload("Items").Order("created_at DESC").Limit(limit).Offset(offset).Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders, "total": total, "page": page})
}

func (h *OrderHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if h.db.First(&order, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}
	var upd struct {
		Status string `json:"status"`
		Notes  string `json:"notes"`
	}
	c.ShouldBindJSON(&upd)
	h.db.Model(&order).Updates(upd)
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetStats(c *gin.Context) {
	var totalOrders, pendingOrders int64
	var totalRevenue float64
	h.db.Model(&models.Order{}).Count(&totalOrders)
	h.db.Model(&models.Order{}).Where("status = ?", "paid").Select("COALESCE(SUM(total),0)").Scan(&totalRevenue)
	h.db.Model(&models.Order{}).Where("status = ?", "pending").Count(&pendingOrders)
	var recent []models.Order
	h.db.Preload("Items").Order("created_at DESC").Limit(5).Find(&recent)
	c.JSON(http.StatusOK, gin.H{
		"total_orders":   totalOrders,
		"total_revenue":  totalRevenue,
		"pending_orders": pendingOrders,
		"recent_orders":  recent,
	})
}
