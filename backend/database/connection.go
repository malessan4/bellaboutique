package database

import (
	"log"

	"bellaboutique/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	log.Println("Conexion a PostgreSQL Neon establecida")
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Error en migraciones:", err)
	}
	log.Println("Migraciones completadas")
}
