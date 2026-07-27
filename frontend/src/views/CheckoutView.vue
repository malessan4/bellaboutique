<template>
  <div class="container section">
    <div class="checkout-layout">
      <!-- Formulario -->
      <div class="checkout-form-wrap">
        <h1 class="display-title" style="margin-bottom:var(--sp-8); font-size: 2.5rem;">Checkout</h1>
        
        <form @submit.prevent="submit" class="checkout-form">
          <div class="form-section">
            <h3 class="form-section-title">Datos de Contacto</h3>
            <div class="form-group">
              <label class="form-label">Email *</label>
              <input type="email" class="form-input" v-model="form.email" required>
            </div>
            <div class="form-group">
              <label class="form-label">Teléfono</label>
              <input type="tel" class="form-input" v-model="form.phone">
            </div>
          </div>

          <div class="form-section">
            <h3 class="form-section-title">Datos de Envío</h3>
            <div class="form-group">
              <label class="form-label">Nombre completo *</label>
              <input type="text" class="form-input" v-model="form.name" required>
            </div>
            <div class="form-group">
              <label class="form-label">Dirección *</label>
              <input type="text" class="form-input" v-model="form.address" required>
            </div>
            <div class="form-row">
              <div class="form-group" style="flex:1">
                <label class="form-label">Ciudad *</label>
                <input type="text" class="form-input" v-model="form.city" required>
              </div>
              <div class="form-group" style="flex:1">
                <label class="form-label">Código Postal *</label>
                <input type="text" class="form-input" v-model="form.zip" required>
              </div>
            </div>
            <div class="form-group">
              <label class="form-label">Provincia *</label>
              <select class="form-select" v-model="form.state" required>
                <option value="">Seleccione...</option>
                <option v-for="p in provincias" :key="p" :value="p">{{ p }}</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Notas del pedido (opcional)</label>
              <textarea class="form-input" v-model="form.notes" rows="3"></textarea>
            </div>
          </div>
          
          <button type="submit" class="btn btn-primary btn-full btn-lg" :disabled="loading" style="margin-top:var(--sp-6)">
            <span v-if="loading">Procesando...</span>
            <span v-else>Pagar con MercadoPago</span>
          </button>
          <p style="text-align:center;font-size:0.8rem;color:var(--c-text-muted);margin-top:var(--sp-3)">
            🔒 Tu pedido es procesado de forma segura con MercadoPago.
          </p>
        </form>
      </div>

      <!-- Resumen -->
      <div class="checkout-summary-wrap">
        <div class="cart-summary glass">
          <h3 class="summary-title" style="font-size:1.25rem">Resumen de compra</h3>
          <div class="summary-items">
            <div class="s-item" v-for="i in cart.items" :key="`${i.id}-${i.size}`">
              <div class="s-img"><img :src="i.image" /><span class="s-qty">{{ i.quantity }}</span></div>
              <div class="s-info">
                <p class="s-name">{{ i.name }}</p>
                <p class="s-var">{{ i.size }} · {{ i.color }}</p>
              </div>
              <div class="s-price">{{ fmt(i.price * i.quantity) }}</div>
            </div>
          </div>
          <div class="divider" style="margin:var(--sp-4) 0"></div>
          <div class="summary-row"><span>Subtotal</span><span>{{ fmt(cart.subtotal) }}</span></div>
          <div class="summary-row"><span>Envío</span><span>{{ cart.shipping===0?'GRATIS':fmt(cart.shipping) }}</span></div>
          <div class="divider" style="margin:var(--sp-4) 0"></div>
          <div class="summary-row summary-total"><span>Total</span><span>{{ fmt(cart.total) }}</span></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, inject } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { ordersApi, paymentsApi } from '@/api'

const cart = useCartStore()
const router = useRouter()
const loading = ref(false)
const toast = inject('toast', null)

const provincias = ['Buenos Aires','Catamarca','Chaco','Chubut','Córdoba','Corrientes','Entre Ríos','Formosa','Jujuy','La Pampa','La Rioja','Mendoza','Misiones','Neuquén','Río Negro','Salta','San Juan','San Luis','Santa Cruz','Santa Fe','Santiago del Estero','Tierra del Fuego','Tucumán']

const form = ref({ email:'', phone:'', name:'', address:'', city:'', zip:'', state:'', notes:'' })

function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p)}

async function submit() {
  if (cart.count === 0) return toast?.error('El carrito está vacío')
  loading.value = true
  try {
    const payload = {
      customer_name: form.value.name,
      customer_email: form.value.email,
      customer_phone: form.value.phone,
      shipping_address: form.value.address,
      city: form.value.city,
      province: form.value.state,
      postal_code: form.value.zip,
      notes: form.value.notes,
      items: cart.items.map(i => ({ product_id: i.id, quantity: i.quantity, size: i.size, color: i.color }))
    }
    const orderRes = await ordersApi.create(payload)
    const orderId = orderRes.data.ID
    
    const payRes = await paymentsApi.create(orderId)
    if(payRes.data.init_point) {
      window.location.href = payRes.data.init_point
    } else {
      toast?.error('El pago ha sido cancelado o cerrado.')
    }
  } catch (e) {
    console.error("Error Checkout:", e.response?.data || e)
    toast?.error('Error al procesar la orden. Intentá nuevamente.')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.checkout-layout { display: grid; grid-template-columns: 1fr 400px; gap: var(--sp-12); align-items: start; }
@media(max-width:992px){.checkout-layout{grid-template-columns:1fr; flex-direction:column-reverse; display:flex;}}

.form-section { margin-bottom: var(--sp-8); }
.form-section-title { font-family: var(--font-display); font-size: 1.5rem; margin-bottom: var(--sp-4); padding-bottom: var(--sp-2); border-bottom: 1px solid var(--c-border); }
.form-row { display: flex; gap: var(--sp-4); }
@media(max-width:480px){.form-row{flex-direction:column;}}

.checkout-summary-wrap { position: sticky; top: 100px; }
.cart-summary { padding: var(--sp-6); border-radius: var(--r-lg); }
.s-item { display: flex; gap: var(--sp-3); align-items: center; margin-bottom: var(--sp-3); }
.s-img { position: relative; width: 64px; height: 64px; border-radius: var(--r-md); background: var(--c-bg-3); flex-shrink: 0; }
.s-img img { width: 100%; height: 100%; object-fit: cover; border-radius: var(--r-md); }
.s-qty { position: absolute; top: -8px; right: -8px; background: var(--c-rose); color: white; width: 20px; height: 20px; border-radius: 50%; font-size: 0.7rem; display: flex; align-items: center; justify-content: center; font-weight: bold; }
.s-info { flex: 1; min-width: 0; }
.s-name { font-size: 0.9rem; font-weight: 500; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.s-var { font-size: 0.8rem; color: var(--c-text-muted); }
.s-price { font-size: 0.95rem; font-weight: 600; }
.summary-row { display: flex; justify-content: space-between; margin-bottom: var(--sp-2); color: var(--c-text-muted); }
.summary-total { font-size: 1.25rem; font-weight: 600; color: var(--c-text); margin-bottom: 0; }
</style>
