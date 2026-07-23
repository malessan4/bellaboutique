# Bella Boutique — E-Commerce

E-commerce de moda femenina y lencería para mercado LATAM/Argentina.

## Tech Stack
- **Frontend**: Vue 3 + Vite + Pinia + Vue Router
- **Backend**: Go + Gin + GORM
- **Base de datos**: PostgreSQL (Neon)
- **Pagos**: MercadoPago Sandbox

## Estructura del Proyecto
```
bellaboutique/
├── backend/    # Go API (puerto 8080)
└── frontend/   # Vue 3 (puerto 5173)
```

## Inicio Rápido

### Backend
```bash
cd backend
# Copiar y configurar variables de entorno
cp .env.example .env
# Editar .env con tus credenciales de MercadoPago sandbox
go run main.go
```

### Frontend
```bash
cd frontend
npm install
cp .env.example .env
npm run dev
```

## MercadoPago Sandbox
Para usar el pago en modo demo, necesitás obtener credenciales sandbox en:
https://www.mercadopago.com.ar/developers/es/docs/checkout-pro/additional-content/your-integrations/credentials

Tarjetas de prueba (Argentina):
- ✅ Aprobado: `4509 9535 6623 3704` — CVV: `123` — Venc: cualquier fecha futura
- ❌ Rechazado: `4000 0000 0000 0002`

## Admin Panel
- URL: http://localhost:5173/admin
- Email: admin@bellaboutique.com
- Password: Bella2024! (configurable en .env)
