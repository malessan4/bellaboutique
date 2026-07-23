<template>
  <div class="container section">
    <div class="catalog-header">
      <h1 class="display-title">Catálogo</h1>
      <button class="btn btn-outline hide-desktop" @click="showFilters=true">Filtros</button>
    </div>

    <div class="catalog-layout">
      <!-- Sidebar / Drawer mobile -->
      <div class="catalog-sidebar" :class="{ 'catalog-sidebar--open': showFilters }">
        <div class="catalog-sidebar__inner">
          <div class="catalog-sidebar__header hide-desktop">
            <h3>Filtros</h3>
            <button @click="showFilters=false" style="color:var(--c-text)">✕</button>
          </div>
          
          <div class="filter-group">
            <h4 class="filter-title">Categorías</h4>
            <label class="filter-radio"><input type="radio" v-model="filters.category" value="" @change="apply"><span>Todas</span></label>
            <label class="filter-radio"><input type="radio" v-model="filters.category" value="lenceria" @change="apply"><span>Lencería</span></label>
            <label class="filter-radio"><input type="radio" v-model="filters.category" value="ropa-casual" @change="apply"><span>Ropa Casual</span></label>
            <label class="filter-radio"><input type="radio" v-model="filters.category" value="pijamas-loungewear" @change="apply"><span>Pijamas & Loungewear</span></label>
            <label class="filter-radio"><input type="radio" v-model="filters.category" value="accesorios" @change="apply"><span>Accesorios</span></label>
          </div>

          <div class="filter-group">
            <h4 class="filter-title">Rango de Precio</h4>
            <div style="display:flex;gap:var(--sp-2)">
              <input type="number" class="form-input" v-model="filters.min" placeholder="Min" @change="apply">
              <input type="number" class="form-input" v-model="filters.max" placeholder="Max" @change="apply">
            </div>
          </div>
          
          <div class="filter-group hide-desktop">
             <button class="btn btn-primary btn-full" @click="showFilters=false">Aplicar</button>
          </div>
        </div>
      </div>
      <div class="overlay hide-desktop" v-if="showFilters" @click="showFilters=false"></div>

      <!-- Main Content -->
      <div class="catalog-main">
        <div class="catalog-topbar">
          <span>{{ total }} productos encontrados</span>
          <select class="form-select" style="width:auto" v-model="filters.sort" @change="apply">
            <option value="">Destacados</option>
            <option value="price_asc">Precio: Menor a Mayor</option>
            <option value="price_desc">Precio: Mayor a Menor</option>
            <option value="name_asc">Nombre: A-Z</option>
          </select>
        </div>

        <div class="products-grid">
          <template v-if="loading">
            <div v-for="i in 8" :key="i" class="skeleton" style="aspect-ratio:3/4"></div>
          </template>
          <template v-else-if="products.length > 0">
            <ProductCard v-for="p in products" :key="p.id" :product="p" />
          </template>
          <div v-else class="empty-state">
            <p>No se encontraron productos.</p>
          </div>
        </div>

        <div class="pagination" v-if="totalPages > 1">
          <button class="btn btn-outline btn-sm" :disabled="page===1" @click="setPage(page-1)">Anterior</button>
          <span>Página {{ page }} de {{ totalPages }}</span>
          <button class="btn btn-outline btn-sm" :disabled="page===totalPages" @click="setPage(page+1)">Siguiente</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ProductCard from '@/components/ui/ProductCard.vue'
import api from '@/api'

const route = useRoute()
const router = useRouter()
const showFilters = ref(false)
const loading = ref(true)
const products = ref([])
const total = ref(0)
const page = ref(1)
const totalPages = ref(1)

const filters = ref({
  category: route.query.category || '',
  search: route.query.search || '',
  min: route.query.min || '',
  max: route.query.max || '',
  sort: route.query.sort || ''
})

async function fetchProducts() {
  loading.value = true
  try {
    const params = new URLSearchParams({ page: page.value, ...filters.value })
    const res = await api.get(`/products?${params.toString()}`)
    products.value = res.data?.data || []
    total.value = res.data?.meta?.total || 0
    totalPages.value = res.data?.meta?.last_page || 1
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

function apply() {
  page.value = 1
  router.push({ query: { ...filters.value } })
}

function setPage(p) {
  page.value = p
  window.scrollTo(0, 0)
  fetchProducts()
}

watch(() => route.query, (q) => {
  filters.value = {
    category: q.category || '',
    search: q.search || '',
    min: q.min || '',
    max: q.max || '',
    sort: q.sort || ''
  }
  fetchProducts()
}, { deep: true })

onMounted(fetchProducts)
</script>

<style scoped>
.catalog-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--sp-8); }
.catalog-layout { display: flex; gap: var(--sp-8); }
.catalog-sidebar { width: 260px; flex-shrink: 0; }
.catalog-main { flex: 1; min-width: 0; }
.catalog-topbar { display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--sp-6); font-size: 0.9rem; color: var(--c-text-muted); }

.filter-group { margin-bottom: var(--sp-6); }
.filter-title { font-size: 1rem; margin-bottom: var(--sp-3); font-family: var(--font-display); }
.filter-radio { display: flex; align-items: center; gap: var(--sp-2); margin-bottom: var(--sp-2); cursor: pointer; color: var(--c-text-muted); font-size: 0.9rem; }
.filter-radio:hover { color: var(--c-text); }
.filter-radio input[type="radio"] { accent-color: var(--c-rose); }

.pagination { display: flex; justify-content: center; align-items: center; gap: var(--sp-4); margin-top: var(--sp-12); color: var(--c-text-muted); font-size: 0.9rem; }
.empty-state { text-align: center; padding: var(--sp-12); color: var(--c-text-muted); grid-column: 1 / -1; }

@media (max-width: 768px) {
  .catalog-sidebar { position: fixed; top: 0; left: 0; width: 300px; height: 100vh; background: var(--c-bg-2); z-index: 300; transform: translateX(-100%); transition: transform var(--t-base); overflow-y: auto; }
  .catalog-sidebar--open { transform: translateX(0); }
  .catalog-sidebar__inner { padding: var(--sp-6); }
  .catalog-sidebar__header { display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--sp-6); padding-bottom: var(--sp-4); border-bottom: 1px solid var(--c-border); }
}
</style>
