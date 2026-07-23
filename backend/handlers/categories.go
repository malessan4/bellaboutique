package handlers

import (
	"net/http"

	"bellaboutique/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct{ db *gorm.DB }

func NewCategoryHandler(db *gorm.DB) *CategoryHandler { return &CategoryHandler{db: db} }

func (h *CategoryHandler) GetAll(c *gin.Context) {
	var categories []models.Category
	h.db.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetBySlug(c *gin.Context) {
	var category models.Category
	if err := h.db.Where("slug = ?", c.Param("slug")).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria no encontrada"})
		return
	}
	c.JSON(http.StatusOK, category)
}
