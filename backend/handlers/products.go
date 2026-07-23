package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"bellaboutique/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct{ db *gorm.DB }

func NewProductHandler(db *gorm.DB) *ProductHandler { return &ProductHandler{db: db} }

func (h *ProductHandler) GetAll(c *gin.Context) {
	var products []models.Product
	query := h.db.Preload("Category")

	if cat := c.Query("category"); cat != "" {
		var category models.Category
		if h.db.Where("slug = ?", cat).First(&category).Error == nil {
			query = query.Where("category_id = ?", category.ID)
		}
	}
	if feat := c.Query("featured"); feat == "true" {
		query = query.Where("featured = ?", true)
	}
	if minP := c.Query("min_price"); minP != "" {
		if v, err := strconv.ParseFloat(minP, 64); err == nil {
			query = query.Where("price >= ?", v)
		}
	}
	if maxP := c.Query("max_price"); maxP != "" {
		if v, err := strconv.ParseFloat(maxP, 64); err == nil {
			query = query.Where("price <= ?", v)
		}
	}
	if search := c.Query("search"); search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var total int64
	query.Model(&models.Product{}).Count(&total)
	query.Limit(limit).Offset(offset).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
		"pages":    (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *ProductHandler) GetFeatured(c *gin.Context) {
	var products []models.Product
	h.db.Preload("Category").Where("featured = ?", true).Limit(8).Find(&products)
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetBySlug(c *gin.Context) {
	var product models.Product
	if err := h.db.Preload("Category").Where("slug = ?", c.Param("slug")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.Slug = slugify(product.Name)
	if err := h.db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear producto"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if h.db.First(&product, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if h.db.Delete(&models.Product{}, id).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}

func slugify(s string) string {
	s = strings.ToLower(s)
	replacer := strings.NewReplacer("á","a","é","e","í","i","ó","o","ú","u","ñ","n","ü","u"," ","-","'","",",","",".","")
	return replacer.Replace(s)
}
