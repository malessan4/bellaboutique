package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	JWTSecret     string
	MPAccessToken string
	MPPublicKey   string
	FrontendURL   string
	Port          string
	AdminEmail    string
	AdminPassword string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Archivo .env no encontrado, usando variables de sistema")
	}
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		JWTSecret:     getEnv("JWT_SECRET", "bella-boutique-secret-jwt-2024"),
		MPAccessToken: getEnv("MP_ACCESS_TOKEN", "TEST-placeholder"),
		MPPublicKey:   getEnv("MP_PUBLIC_KEY", "TEST-placeholder"),
		FrontendURL:   getEnv("FRONTEND_URL", "http://localhost:5173"),
		Port:          getEnv("PORT", "8080"),
		AdminEmail:    getEnv("ADMIN_EMAIL", "admin@bellaboutique.com"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "Bella2024!"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
