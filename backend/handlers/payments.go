package handlers

import (
	"log"
	"net/http"
	"strconv"

	"bellaboutique/config"
	"bellaboutique/models"
	"bellaboutique/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewPaymentHandler(db *gorm.DB, cfg *config.Config) *PaymentHandler {
	return &PaymentHandler{db: db, cfg: cfg}
}

func (h *PaymentHandler) Create(c *gin.Context) {
	var req struct {
		OrderID uint `json:"order_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var order models.Order
	if h.db.Preload("Items").First(&order, req.OrderID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}
	pref, err := services.CreateMPPreference(&order, order.Items, h.cfg.MPAccessToken, h.cfg.FrontendURL)
	if err != nil {
		log.Println("ERROR MercadoPago:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error MercadoPago: " + err.Error()})
		return
	}
	h.db.Model(&order).Update("mp_preference_id", pref.ID)
	c.JSON(http.StatusOK, gin.H{
		"preference_id":      pref.ID,
		"init_point":         pref.InitPoint,
		"sandbox_init_point": pref.SandboxInitPoint,
	})
}

func (h *PaymentHandler) Webhook(c *gin.Context) {
	var payload struct {
		Action string `json:"action"`
		Data   struct {
			ID string `json:"id"`
		} `json:"data"`
		Type string `json:"type"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload invalido"})
		return
	}
	if payload.Type == "payment" {
		mpPayment, err := services.GetMPPayment(payload.Data.ID, h.cfg.MPAccessToken)
		if err == nil {
			orderID, _ := strconv.Atoi(mpPayment.ExternalReference)
			if orderID > 0 {
				status := map[string]string{
					"approved":  "paid",
					"rejected":  "cancelled",
					"cancelled": "cancelled",
				}[mpPayment.Status]
				if status == "" {
					status = "pending"
				}
				h.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(map[string]interface{}{
					"status":        status,
					"mp_payment_id": payload.Data.ID,
				})

				// Descontar stock si el pago fue aprobado
				if status == "paid" {
					var orderItems []models.OrderItem
					h.db.Where("order_id = ?", orderID).Find(&orderItems)
					for _, item := range orderItems {
						h.db.Exec("UPDATE products SET stock = stock - ? WHERE id = ? AND stock >= ?", item.Quantity, item.ProductID, item.Quantity)
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
