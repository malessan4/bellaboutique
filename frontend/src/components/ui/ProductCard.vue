<template>
  <article class="pc" @click="go">
    <div class="pc__img-wrap">
      <img v-if="!imgErr" :src="product.images?.[0]" :alt="product.name" class="pc__img" loading="lazy" @error="imgErr=true" />
      <div v-else class="pc__img-ph">♡</div>
      <div class="pc__badges">
        <span v-if="product.sale_price" class="badge badge-sale">OFERTA</span>
      </div>
      <div class="pc__overlay">
        <button class="btn btn-primary btn-sm" @click.stop="quickAdd">+ Agregar</button>
      </div>
    </div>
    <div class="pc__info">
      <p class="pc__cat">{{ product.category?.name }}</p>
      <h3 class="pc__name">{{ product.name }}</h3>
      <div class="pc__price">
        <span v-if="product.sale_price" class="price price-sale">{{ fmt(product.sale_price) }}</span>
        <span v-else class="price">{{ fmt(product.price) }}</span>
        <span v-if="product.sale_price" class="price-original">{{ fmt(product.price) }}</span>
      </div>
    </div>
  </article>
</template>

<script setup>
import { ref, inject } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'

const props = defineProps({ product: { type: Object, required: true } })
const router = useRouter()
const cart = useCartStore()
const toast = inject('toast', null)
const imgErr = ref(false)

function go() { router.push(`/producto/${props.product.slug}`) }
function quickAdd() {
  cart.addItem(props.product, props.product.sizes?.[0] || '', props.product.colors?.[0] || '')
  toast?.success(`${props.product.name} agregado al carrito`)
}
function fmt(p) { return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p) }
</script>

<style scoped>
.pc{cursor:pointer;border-radius:var(--r-lg);overflow:hidden;background:var(--c-card);border:1px solid var(--c-border);transition:all var(--t-base);}
.pc:hover{border-color:var(--c-border-rose);transform:translateY(-4px);box-shadow:var(--shadow-card);}
.pc__img-wrap{position:relative;aspect-ratio:3/4;overflow:hidden;background:var(--c-bg-3);}
.pc__img{width:100%;height:100%;object-fit:cover;transition:transform 0.5s ease;}
.pc:hover .pc__img{transform:scale(1.05);}
.pc__img-ph{width:100%;height:100%;display:flex;align-items:center;justify-content:center;font-size:3rem;color:var(--c-border-rose);}
.pc__badges{position:absolute;top:var(--sp-3);left:var(--sp-3);display:flex;flex-direction:column;gap:var(--sp-1);}
.pc__overlay{position:absolute;bottom:0;left:0;right:0;padding:var(--sp-4);background:linear-gradient(0deg,rgba(13,10,11,0.9),transparent);transform:translateY(100%);transition:transform var(--t-base);display:flex;justify-content:center;}
.pc:hover .pc__overlay{transform:translateY(0);}
.pc__info{padding:var(--sp-4);}
.pc__cat{font-size:0.7rem;color:var(--c-rose);letter-spacing:0.08em;text-transform:uppercase;margin-bottom:var(--sp-1);}
.pc__name{font-family:var(--font-display);font-size:0.95rem;font-weight:500;margin-bottom:var(--sp-2);line-height:1.3;}
.pc__price{display:flex;align-items:center;gap:var(--sp-2);}
</style>
