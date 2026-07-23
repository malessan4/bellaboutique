package main

import (
	"log"

	"bellaboutique/config"
	"bellaboutique/database"
	"bellaboutique/routes"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	database.AutoMigrate(db)
	database.Seed(db, cfg)

	r := routes.Setup(db, cfg)
	log.Printf("Bella Boutique API iniciado en puerto %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
