<template>
  <div class="container section" v-if="product">
    <nav class="breadcrumb">
      <router-link to="/">Inicio</router-link> >
      <router-link :to="`/catalogo?category=${product.category?.slug}`">{{ product.category?.name }}</router-link> >
      <span>{{ product.name }}</span>
    </nav>

    <div class="product-layout">
      <!-- Gallery -->
      <div class="product-gallery">
        <div class="product-img-main">
          <img :src="mainImage" :alt="product.name" />
          <span v-if="product.sale_price" class="badge badge-sale product-badge">OFERTA</span>
        </div>
        <div class="product-thumbnails" v-if="product.images?.length > 1">
          <button v-for="(img, idx) in product.images" :key="idx" class="thumbnail" :class="{active: mainImage === img}" @click="mainImage = img">
            <img :src="img" alt="Thumbnail" />
          </button>
        </div>
      </div>

      <!-- Info -->
      <div class="product-info">
        <h1 class="display-title product-title">{{ product.name }}</h1>
        <div class="product-price">
          <span v-if="product.sale_price" class="price price-sale" style="font-size:2rem">{{ fmt(product.sale_price) }}</span>
          <span v-else class="price" style="font-size:2rem">{{ fmt(product.price) }}</span>
          <span v-if="product.sale_price" class="price-original" style="font-size:1.2rem">{{ fmt(product.price) }}</span>
        </div>

        <p class="product-desc">{{ product.description }}</p>

        <!-- Variants -->
        <div class="product-options" v-if="product.sizes?.length">
          <span class="option-label">Talle</span>
          <div class="size-chips">
            <button v-for="size in product.sizes" :key="size" class="size-chip" :class="{active: selectedSize === size}" @click="selectedSize = size">{{ size }}</button>
          </div>
        </div>

        <div class="product-options" v-if="product.colors?.length">
          <span class="option-label">Color</span>
          <div class="color-chips">
            <button v-for="color in product.colors" :key="color" class="color-chip" :class="{active: selectedColor === color}" :style="{ backgroundColor: cssColor(color) }" @click="selectedColor = color" :title="color"></button>
          </div>
        </div>

        <!-- Add to Cart -->
        <div class="product-actions">
          <div class="qty-selector">
            <button @click="qty = Math.max(1, qty - 1)">-</button>
            <span>{{ qty }}</span>
            <button @click="qty++">+</button>
          </div>
          <button class="btn btn-primary btn-lg" style="flex:1" @click="addToCart">Agregar al carrito</button>
        </div>

        <!-- Extras -->
        <div class="product-extras">
          <div class="extra-item">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="1" y="3" width="15" height="13"/><polygon points="16 8 20 8 23 11 23 16 16 16 16 8"/><circle cx="5.5" cy="18.5" r="2.5"/><circle cx="18.5" cy="18.5" r="2.5"/></svg>
            Envío gratis en compras mayores a $30.000
          </div>
          <div class="extra-item">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>
            Devoluciones dentro de los 30 días
          </div>
        </div>
      </div>
    </div>

    <!-- Related -->
    <div class="related-section" v-if="related.length">
      <h2 class="section-title" style="text-align:center;margin-bottom:var(--sp-8)">También te puede gustar</h2>
      <div class="products-grid">
        <ProductCard v-for="p in related" :key="p.id" :product="p" />
      </div>
    </div>
  </div>
  <div v-else-if="loading" class="container section" style="text-align:center;padding:var(--sp-24)">
    Cargando...
  </div>
</template>

<script setup>
import { ref, onMounted, inject, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { productsApi } from '@/api'
import ProductCard from '@/components/ui/ProductCard.vue'

const route = useRoute()
const cart = useCartStore()
const toast = inject('toast')
const product = ref(null)
const related = ref([])
const loading = ref(true)

const mainImage = ref('')
const selectedSize = ref('')
const selectedColor = ref('')
const qty = ref(1)

function cssColor(c) {
  const map = { 'blanco': '#ffffff', 'negro': '#000000', 'rojo': '#ff0000', 'rosa': '#ffc0cb', 'nude': '#e3bc9a' }
  return map[c.toLowerCase()] || c
}
function fmt(p) { return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p) }

async function loadProduct() {
  loading.value = true
  try {
    const res = await productsApi.getBySlug(route.params.slug)
    product.value = res.data
    mainImage.value = product.value.images?.[0] || ''
    selectedSize.value = product.value.sizes?.[0] || ''
    selectedColor.value = product.value.colors?.[0] || ''
    
    // Load related
    const relRes = await productsApi.getAll({ category: product.value.category?.slug, limit: 4 })
    related.value = (relRes.data.products || []).filter(p => p.id !== product.value.id).slice(0,4)
  } catch(e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

function addToCart() {
  if (product.value.sizes?.length && !selectedSize.value) return toast.error('Selecciona un talle')
  if (product.value.colors?.length && !selectedColor.value) return toast.error('Selecciona un color')
  
  cart.addItem(product.value, selectedSize.value, selectedColor.value, qty.value)
  toast.success('Producto agregado al carrito')
}

watch(() => route.params.slug, loadProduct)
onMounted(loadProduct)
</script>

<style scoped>
.breadcrumb { font-size: 0.85rem; color: var(--c-text-muted); margin-bottom: var(--sp-8); display: flex; gap: var(--sp-2); align-items: center; }
.breadcrumb a:hover { color: var(--c-rose); }

.product-layout { display: grid; grid-template-columns: 1fr 1fr; gap: var(--sp-12); margin-bottom: var(--sp-24); }
@media (max-width: 768px) { .product-layout { grid-template-columns: 1fr; gap: var(--sp-8); } }

.product-img-main { position: relative; aspect-ratio: 3/4; border-radius: var(--r-lg); overflow: hidden; background: var(--c-bg-3); }
.product-img-main img { width: 100%; height: 100%; object-fit: cover; }
.product-badge { position: absolute; top: var(--sp-4); left: var(--sp-4); }

.product-thumbnails { display: flex; gap: var(--sp-3); margin-top: var(--sp-3); overflow-x: auto; padding-bottom: var(--sp-2); }
.thumbnail { width: 80px; height: 100px; flex-shrink: 0; border-radius: var(--r-md); overflow: hidden; border: 2px solid transparent; transition: border-color var(--t-fast); }
.thumbnail img { width: 100%; height: 100%; object-fit: cover; }
.thumbnail.active { border-color: var(--c-rose); }

.product-title { font-size: 2.5rem; margin-bottom: var(--sp-4); }
.product-price { display: flex; align-items: center; gap: var(--sp-3); margin-bottom: var(--sp-6); }
.product-desc { color: var(--c-text-muted); line-height: 1.8; margin-bottom: var(--sp-8); }

.product-options { margin-bottom: var(--sp-6); }
.option-label { display: block; font-size: 0.85rem; text-transform: uppercase; letter-spacing: 0.05em; color: var(--c-text-muted); margin-bottom: var(--sp-3); }

.size-chips { display: flex; flex-wrap: wrap; gap: var(--sp-3); }
.size-chip { padding: 0.5rem 1.5rem; border: 1px solid var(--c-border); border-radius: var(--r-md); background: transparent; color: var(--c-text); font-weight: 500; transition: all var(--t-fast); }
.size-chip:hover { border-color: var(--c-rose-light); }
.size-chip.active { background: var(--c-rose); border-color: var(--c-rose); color: white; }

.color-chips { display: flex; gap: var(--sp-3); }
.color-chip { width: 36px; height: 36px; border-radius: var(--r-full); border: 2px solid var(--c-border); cursor: pointer; transition: transform var(--t-fast); }
.color-chip:hover { transform: scale(1.1); }
.color-chip.active { box-shadow: 0 0 0 2px var(--c-bg), 0 0 0 4px var(--c-rose); border-color: transparent; }

.product-actions { display: flex; gap: var(--sp-4); margin-bottom: var(--sp-8); margin-top: var(--sp-8); }
.qty-selector { display: flex; align-items: center; border: 1px solid var(--c-border); border-radius: var(--r-full); padding: 0.25rem; }
.qty-selector button { width: 40px; height: 40px; display: flex; align-items: center; justify-content: center; color: var(--c-text); font-size: 1.2rem; }
.qty-selector span { width: 40px; text-align: center; font-weight: 500; }

.product-extras { display: flex; flex-direction: column; gap: var(--sp-3); padding-top: var(--sp-6); border-top: 1px solid var(--c-border); }
.extra-item { display: flex; align-items: center; gap: var(--sp-3); color: var(--c-text-muted); font-size: 0.9rem; }
.extra-item svg { color: var(--c-rose); }
</style>
