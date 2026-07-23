<template>
  <Teleport to="body">
    <div class="overlay" v-if="cart.isOpen" @click="cart.isOpen=false"></div>
    <div class="cd" :class="{open:cart.isOpen}">
      <div class="cd__head">
        <div><h3 class="cd__title">Mi Carrito</h3><p class="cd__sub">{{ cart.count }} artículo{{ cart.count!==1?'s':'' }}</p></div>
        <button class="cd__close" @click="cart.isOpen=false">✕</button>
      </div>

      <div class="cd__banner" v-if="cart.count>0">
        <template v-if="cart.subtotal<30000">
          <p>🚚 Te faltan <strong>{{ fmt(30000-cart.subtotal) }}</strong> para envío gratis</p>
          <div class="cd__prog"><div class="cd__prog-bar" :style="{width:Math.min(100,(cart.subtotal/30000)*100)+'%'}"></div></div>
        </template>
        <p v-else style="color:#48c78e">🎉 ¡Tenés envío gratis!</p>
      </div>

      <div class="cd__items" v-if="cart.count>0">
        <div class="ci" v-for="i in cart.items" :key="`${i.id}-${i.size}-${i.color}`">
          <div class="ci__img"><img :src="i.image" :alt="i.name" loading="lazy" /></div>
          <div class="ci__info">
            <p class="ci__name">{{ i.name }}</p>
            <p class="ci__var">{{ i.size }} · {{ i.color }}</p>
            <p class="ci__price">{{ fmt(i.price) }}</p>
          </div>
          <div class="ci__qty">
            <button @click="cart.updateQuantity(i.id,i.size,i.color,i.quantity-1)">−</button>
            <span>{{ i.quantity }}</span>
            <button @click="cart.updateQuantity(i.id,i.size,i.color,i.quantity+1)">+</button>
          </div>
          <button class="ci__del" @click="cart.removeItem(i.id,i.size,i.color)" title="Eliminar">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14H6L5 6"/></svg>
          </button>
        </div>
      </div>

      <div class="cd__empty" v-else>
        <p style="font-size:2.5rem">🛍️</p>
        <p>Tu carrito está vacío</p>
        <router-link to="/catalogo" class="btn btn-primary" @click="cart.isOpen=false">Ver productos</router-link>
      </div>

      <div class="cd__foot" v-if="cart.count>0">
        <div class="cd__totals">
          <div class="cd__row"><span>Subtotal</span><span>{{ fmt(cart.subtotal) }}</span></div>
          <div class="cd__row"><span>Envío</span><span :style="cart.shipping===0?'color:#48c78e':''" >{{ cart.shipping===0?'GRATIS':fmt(cart.shipping) }}</span></div>
          <div class="cd__row cd__row--total"><span>Total</span><span>{{ fmt(cart.total) }}</span></div>
        </div>
        <router-link to="/checkout" class="btn btn-primary btn-full" @click="cart.isOpen=false">Finalizar compra →</router-link>
        <router-link to="/carrito" class="btn btn-outline btn-full" style="margin-top:0.5rem" @click="cart.isOpen=false">Ver carrito</router-link>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { useCartStore } from '@/stores/cart'
const cart = useCartStore()
function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p)}
</script>

<style scoped>
.cd{position:fixed;top:0;right:0;width:380px;height:100vh;background:var(--c-bg-2);border-left:1px solid var(--c-border);z-index:300;transform:translateX(100%);transition:transform var(--t-base);display:flex;flex-direction:column;overflow:hidden;}
.cd.open{transform:translateX(0);}
@media(max-width:420px){.cd{width:100%;}}
.cd__head{display:flex;align-items:flex-start;justify-content:space-between;padding:var(--sp-6);border-bottom:1px solid var(--c-border);flex-shrink:0;}
.cd__title{font-family:var(--font-display);font-size:1.25rem;}
.cd__sub{font-size:0.8rem;color:var(--c-text-muted);margin-top:2px;}
.cd__close{color:var(--c-text-muted);font-size:1.1rem;padding:var(--sp-1);transition:color var(--t-fast);}
.cd__close:hover{color:var(--c-text);}
.cd__banner{padding:var(--sp-3) var(--sp-6);background:rgba(201,132,122,0.08);border-bottom:1px solid var(--c-border-rose);font-size:0.8rem;color:var(--c-text-muted);flex-shrink:0;}
.cd__prog{height:3px;background:var(--c-border);border-radius:2px;margin-top:var(--sp-2);overflow:hidden;}
.cd__prog-bar{height:100%;background:linear-gradient(90deg,var(--c-rose),var(--c-gold));border-radius:2px;transition:width 0.5s ease;}
.cd__items{flex:1;overflow-y:auto;padding:var(--sp-4) var(--sp-6);display:flex;flex-direction:column;gap:var(--sp-4);}
.ci{display:flex;align-items:center;gap:var(--sp-3);padding-bottom:var(--sp-4);border-bottom:1px solid var(--c-border);}
.ci__img{width:64px;height:80px;border-radius:var(--r-md);overflow:hidden;flex-shrink:0;background:var(--c-bg-3);}
.ci__img img{width:100%;height:100%;object-fit:cover;}
.ci__info{flex:1;min-width:0;}
.ci__name{font-size:0.85rem;font-weight:500;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;}
.ci__var{font-size:0.75rem;color:var(--c-text-muted);margin-top:2px;}
.ci__price{font-size:0.9rem;font-weight:600;color:var(--c-rose);margin-top:var(--sp-1);}
.ci__qty{display:flex;align-items:center;gap:var(--sp-2);flex-shrink:0;}
.ci__qty button{width:24px;height:24px;border-radius:var(--r-full);border:1px solid var(--c-border);color:var(--c-text-muted);display:flex;align-items:center;justify-content:center;font-size:1rem;transition:all var(--t-fast);}
.ci__qty button:hover{border-color:var(--c-rose);color:var(--c-rose);}
.ci__qty span{font-size:0.85rem;width:20px;text-align:center;}
.ci__del{color:var(--c-text-subtle);padding:var(--sp-1);transition:color var(--t-fast);}
.ci__del:hover{color:#ff5b5b;}
.cd__empty{flex:1;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:var(--sp-4);padding:var(--sp-8);color:var(--c-text-muted);}
.cd__foot{padding:var(--sp-6);border-top:1px solid var(--c-border);flex-shrink:0;}
.cd__totals{display:flex;flex-direction:column;gap:var(--sp-2);margin-bottom:var(--sp-4);}
.cd__row{display:flex;justify-content:space-between;font-size:0.9rem;color:var(--c-text-muted);}
.cd__row--total{font-size:1.1rem;font-weight:600;color:var(--c-text);padding-top:var(--sp-2);border-top:1px solid var(--c-border);margin-top:var(--sp-2);}
</style>
