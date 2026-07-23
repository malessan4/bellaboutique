package routes

import (
	"bellaboutique/config"
	"bellaboutique/handlers"
	"bellaboutique/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL, "http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	productH := handlers.NewProductHandler(db)
	categoryH := handlers.NewCategoryHandler(db)
	orderH := handlers.NewOrderHandler(db, cfg)
	authH := handlers.NewAuthHandler(db, cfg)
	paymentH := handlers.NewPaymentHandler(db, cfg)

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "shop": "Bella Boutique"})
		})

		api.GET("/categories", categoryH.GetAll)
		api.GET("/categories/:slug", categoryH.GetBySlug)

		api.GET("/products", productH.GetAll)
		api.GET("/products/featured", productH.GetFeatured)
		api.GET("/products/:slug", productH.GetBySlug)

		api.POST("/orders", orderH.Create)
		api.GET("/orders/:id", orderH.GetByID)

		api.POST("/payments/create", paymentH.Create)
		api.POST("/payments/webhook", paymentH.Webhook)

		api.POST("/auth/login", authH.Login)

		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(cfg.JWTSecret))
		{
			admin.POST("/products", productH.Create)
			admin.PUT("/products/:id", productH.Update)
			admin.DELETE("/products/:id", productH.Delete)
			admin.GET("/orders", orderH.GetAll)
			admin.PUT("/orders/:id", orderH.Update)
			admin.GET("/stats", orderH.GetStats)
		}
	}
	return r
}
