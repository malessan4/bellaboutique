<template>
  <div class="success-page">
    <div class="confetti-container">
      <div v-for="i in 50" :key="i" class="confetti" :style="confettiStyle(i)"></div>
    </div>
    
    <div class="success-card glass">
      <div class="icon-wrap">
        <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 11-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>
      </div>
      
      <h1 class="display-title" style="font-size:2.5rem; margin-bottom:var(--sp-2)">¡Gracias por tu compra!</h1>
      <p style="color:var(--c-text-muted); font-size:1.1rem; margin-bottom:var(--sp-6)">Tu pedido ha sido procesado exitosamente.</p>
      
      <div class="order-info">
        <p>Número de orden: <strong>#{{ orderId }}</strong></p>
        <p style="font-size:0.9rem; margin-top:var(--sp-2)">Recibirás un email con los detalles de tu pedido a la brevedad.</p>
      </div>
      
      <div class="actions">
        <router-link to="/catalogo" class="btn btn-primary btn-lg">Seguir comprando →</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useCartStore } from '@/stores/cart'

const route = useRoute()
const cart = useCartStore()
const orderId = route.query.order || '1000'

onMounted(() => {
  cart.clear()
})

function confettiStyle(i) {
  const left = Math.random() * 100
  const animDelay = Math.random() * 3
  const color = ['#c9847a', '#e2c992', '#f0c4cc', '#ffffff'][Math.floor(Math.random()*4)]
  return `left:${left}%; animation-delay:${animDelay}s; background-color:${color};`
}
</script>

<style scoped>
.success-page { min-height: 80vh; display: flex; align-items: center; justify-content: center; padding: var(--sp-6); position: relative; overflow: hidden; }
.success-card { max-width: 600px; width: 100%; text-align: center; padding: var(--sp-12); border-radius: var(--r-xl); position: relative; z-index: 10; animation: fadeUp 0.6s ease; }
.icon-wrap { width: 80px; height: 80px; border-radius: 50%; background: rgba(72,199,142,0.1); color: #48c78e; display: flex; align-items: center; justify-content: center; margin: 0 auto var(--sp-6); animation: scaleIn 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) 0.2s both; }
.check-icon { width: 40px; height: 40px; }

.order-info { background: rgba(0,0,0,0.2); padding: var(--sp-6); border-radius: var(--r-lg); margin-bottom: var(--sp-8); }
.actions { display: flex; justify-content: center; }

@keyframes scaleIn { from { transform: scale(0); } to { transform: scale(1); } }

.confetti-container { position: absolute; inset: 0; z-index: 1; pointer-events: none; }
.confetti { position: absolute; top: -10px; width: 10px; height: 10px; opacity: 0; animation: fall 4s linear infinite; }
@keyframes fall {
  0% { transform: translateY(-10vh) rotate(0deg); opacity: 1; }
  100% { transform: translateY(110vh) rotate(720deg); opacity: 0; }
}
</style>
