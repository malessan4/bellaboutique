<template>
  <div class="container section">
    <h1 class="display-title" style="margin-bottom:var(--sp-8)">Mi Carrito</h1>

    <div v-if="cart.count === 0" class="empty-cart">
      <div class="empty-icon">🛍️</div>
      <h2>Tu carrito está vacío</h2>
      <p>¡Descubrí nuestra colección y encontrá tu próximo look favorito!</p>
      <router-link to="/catalogo" class="btn btn-primary btn-lg" style="margin-top:var(--sp-4)">Explorar Catálogo</router-link>
    </div>

    <div v-else class="cart-layout">
      <!-- Items -->
      <div class="cart-items">
        <div class="cart-table-header hide-mobile">
          <span>Producto</span>
          <span>Cantidad</span>
          <span>Subtotal</span>
          <span></span>
        </div>
        
        <div class="cart-item" v-for="i in cart.items" :key="`${i.id}-${i.size}-${i.color}`">
          <div class="item-prod">
            <img :src="i.image" :alt="i.name" />
            <div class="item-info">
              <h3>{{ i.name }}</h3>
              <p class="item-var">{{ i.size }} · {{ i.color }}</p>
              <p class="item-price hide-desktop">{{ fmt(i.price) }}</p>
            </div>
          </div>
          <div class="item-qty">
            <div class="qty-control">
              <button @click="cart.updateQuantity(i.id, i.size, i.color, i.quantity - 1)">-</button>
              <span>{{ i.quantity }}</span>
              <button @click="cart.updateQuantity(i.id, i.size, i.color, i.quantity + 1)">+</button>
            </div>
          </div>
          <div class="item-subtotal hide-mobile">{{ fmt(i.price * i.quantity) }}</div>
          <div class="item-action">
            <button class="btn-ghost" @click="cart.removeItem(i.id, i.size, i.color)">Eliminar</button>
          </div>
        </div>
      </div>

      <!-- Summary -->
      <div class="cart-summary-wrap">
        <div class="cart-summary glass">
          <h3 class="summary-title">Resumen de compra</h3>
          
          <div class="summary-row">
            <span>Subtotal ({{ cart.count }} ítems)</span>
            <span>{{ fmt(cart.subtotal) }}</span>
          </div>
          <div class="summary-row">
            <span>Envío</span>
            <span :style="cart.shipping===0?'color:#48c78e':''">{{ cart.shipping === 0 ? 'GRATIS' : fmt(cart.shipping) }}</span>
          </div>
          
          <div class="divider"></div>
          
          <div class="summary-row summary-total">
            <span>Total</span>
            <span>{{ fmt(cart.total) }}</span>
          </div>

          <router-link to="/checkout" class="btn btn-primary btn-full btn-lg" style="margin-top:var(--sp-6)">Ir al checkout</router-link>
          
          <div class="shipping-progress" style="margin-top:var(--sp-6)">
            <template v-if="cart.subtotal < 30000">
              <p style="font-size:0.8rem;text-align:center;margin-bottom:var(--sp-2);color:var(--c-text-muted)">Te faltan <strong>{{ fmt(30000 - cart.subtotal) }}</strong> para envío gratis</p>
              <div class="prog-bg"><div class="prog-fill" :style="{width: (cart.subtotal/30000)*100 + '%'}"></div></div>
            </template>
            <p v-else style="font-size:0.85rem;text-align:center;color:#48c78e;font-weight:500">🎉 ¡Tenés envío gratis!</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '@/stores/cart'
const cart = useCartStore()
function fmt(p) { return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p) }
</script>

<style scoped>
.empty-cart { text-align: center; padding: var(--sp-16) 0; }
.empty-icon { font-size: 4rem; margin-bottom: var(--sp-4); }
.empty-cart p { color: var(--c-text-muted); margin-top: var(--sp-2); }

.cart-layout { display: grid; grid-template-columns: 1fr 380px; gap: var(--sp-12); align-items: start; }
@media (max-width: 992px) { .cart-layout { grid-template-columns: 1fr; } }

.cart-table-header { display: grid; grid-template-columns: 2fr 1fr 1fr auto; padding-bottom: var(--sp-4); border-bottom: 1px solid var(--c-border); color: var(--c-text-muted); font-size: 0.85rem; text-transform: uppercase; letter-spacing: 0.05em; }
.cart-item { display: grid; grid-template-columns: 2fr 1fr 1fr auto; align-items: center; padding: var(--sp-6) 0; border-bottom: 1px solid var(--c-border); }
@media (max-width: 768px) { .cart-item { grid-template-columns: 1fr; gap: var(--sp-4); } .item-action { text-align: right; } }

.item-prod { display: flex; gap: var(--sp-4); align-items: center; }
.item-prod img { width: 80px; height: 100px; object-fit: cover; border-radius: var(--r-md); background: var(--c-bg-3); }
.item-info h3 { font-size: 1rem; font-weight: 500; margin-bottom: var(--sp-1); }
.item-var { font-size: 0.85rem; color: var(--c-text-muted); }

.qty-control { display: inline-flex; align-items: center; border: 1px solid var(--c-border); border-radius: var(--r-full); padding: 0.25rem; }
.qty-control button { width: 32px; height: 32px; font-size: 1.1rem; color: var(--c-text); }
.qty-control span { width: 32px; text-align: center; font-size: 0.9rem; }

.item-subtotal { font-weight: 600; font-size: 1.1rem; }

.cart-summary-wrap { position: sticky; top: 100px; }
.cart-summary { padding: var(--sp-6); border-radius: var(--r-lg); }
.summary-title { font-family: var(--font-display); font-size: 1.5rem; margin-bottom: var(--sp-6); }
.summary-row { display: flex; justify-content: space-between; margin-bottom: var(--sp-3); color: var(--c-text-muted); }
.summary-total { font-size: 1.25rem; font-weight: 600; color: var(--c-text); margin-bottom: 0; }

.prog-bg { height: 4px; background: var(--c-border); border-radius: 2px; overflow: hidden; }
.prog-fill { height: 100%; background: linear-gradient(90deg, var(--c-rose), var(--c-gold)); transition: width 0.3s ease; }
</style>
