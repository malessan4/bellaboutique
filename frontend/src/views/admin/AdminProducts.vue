<template>
  <div class="admin-layout">
    <AdminSidebar />
    <main class="admin-main">
      <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:var(--sp-8)">
        <h1 class="display-title" style="font-size:2rem;">Productos</h1>
        <button class="btn btn-primary" @click="openModal(null)">+ Nuevo Producto</button>
      </div>

      <div class="table-wrap glass">
        <table class="admin-table">
          <thead>
            <tr>
              <th>Nombre</th>
              <th>Categoría</th>
              <th>Precio</th>
              <th>Stock</th>
              <th>Destacado</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in products" :key="p.id">
              <td>{{ p.name }}</td>
              <td>{{ p.category?.name || '-' }}</td>
              <td>{{ fmt(p.price) }}</td>
              <td>{{ p.stock }}</td>
              <td>{{ p.is_featured ? 'Sí' : 'No' }}</td>
              <td>
                <button class="btn-ghost" style="padding:0.25rem" @click="openModal(p)">Editar</button>
                <button class="btn-ghost" style="padding:0.25rem;color:#ff5b5b" @click="del(p.id)">Eliminar</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>

    <!-- Modal -->
    <div class="overlay" v-if="modalOpen">
      <div class="modal glass">
        <h2 style="margin-bottom:var(--sp-6)">{{ form.id ? 'Editar Producto' : 'Nuevo Producto' }}</h2>
        <form @submit.prevent="save">
          <div class="form-group"><label class="form-label">Nombre</label><input type="text" class="form-input" v-model="form.name" required></div>
          <div class="form-group"><label class="form-label">Descripción</label><textarea class="form-input" v-model="form.description" required></textarea></div>
          <div style="display:flex;gap:var(--sp-4)">
            <div class="form-group" style="flex:1"><label class="form-label">Precio</label><input type="number" class="form-input" v-model="form.price" required></div>
            <div class="form-group" style="flex:1"><label class="form-label">Precio Oferta</label><input type="number" class="form-input" v-model="form.sale_price"></div>
          </div>
          <div class="form-group">
            <label class="form-label">Categoría</label>
            <select class="form-select" v-model="form.category_id" required>
              <option v-for="c in categories" :key="c.id" :value="c.id">{{ c.name }}</option>
            </select>
          </div>
          <div class="form-group"><label class="form-label">Imágenes (URLs separadas por coma)</label><input type="text" class="form-input" v-model="form.imagesStr"></div>
          <div style="display:flex;gap:var(--sp-4)">
            <div class="form-group" style="flex:1"><label class="form-label">Talles (coma)</label><input type="text" class="form-input" v-model="form.sizesStr"></div>
            <div class="form-group" style="flex:1"><label class="form-label">Colores (coma)</label><input type="text" class="form-input" v-model="form.colorsStr"></div>
          </div>
          <div class="form-group"><label class="form-label">Stock</label><input type="number" class="form-input" v-model="form.stock" required></div>
          <div class="form-group" style="flex-direction:row;align-items:center;">
            <input type="checkbox" v-model="form.is_featured" id="feat"><label for="feat">Destacado</label>
          </div>
          
          <div style="display:flex;justify-content:flex-end;gap:var(--sp-4);margin-top:var(--sp-6)">
            <button type="button" class="btn btn-ghost" @click="modalOpen=false">Cancelar</button>
            <button type="submit" class="btn btn-primary">Guardar</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AdminSidebar from '@/components/admin/AdminSidebar.vue'
import api from '@/api'

const products = ref([])
const categories = ref([])
const modalOpen = ref(false)
const form = ref({})

function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p)}

async function load() {
  try {
    const res = await api.get('/products')
    products.value = res.data.data
    const cRes = await api.get('/categories')
    categories.value = cRes.data
  } catch(e) { console.error(e) }
}

function openModal(p) {
  if (p) {
    form.value = { ...p, imagesStr: p.images?.join(',')||'', sizesStr: p.sizes?.join(',')||'', colorsStr: p.colors?.join(',')||'' }
  } else {
    form.value = { name:'', description:'', price:0, stock:0, is_featured:false, imagesStr:'', sizesStr:'', colorsStr:'' }
  }
  modalOpen.value = true
}

async function save() {
  try {
    const payload = { ...form.value, 
      images: form.value.imagesStr?form.value.imagesStr.split(',').map(s=>s.trim()):[],
      sizes: form.value.sizesStr?form.value.sizesStr.split(',').map(s=>s.trim()):[],
      colors: form.value.colorsStr?form.value.colorsStr.split(',').map(s=>s.trim()):[]
    }
    if (form.value.id) await api.put(`/products/${form.value.id}`, payload)
    else await api.post('/products', payload)
    modalOpen.value = false
    load()
  } catch(e) { alert('Error al guardar') }
}

async function del(id) {
  if(confirm('¿Eliminar producto?')) {
    await api.delete(`/products/${id}`)
    load()
  }
}

onMounted(load)
</script>

<style scoped>
.admin-layout { display: flex; min-height: 100vh; }
.admin-main { flex: 1; padding: var(--sp-8); overflow-y: auto; background: var(--c-bg); }
.table-wrap { border-radius: var(--r-lg); overflow-x: auto; }
.admin-table { width: 100%; border-collapse: collapse; }
.admin-table th, .admin-table td { padding: var(--sp-4); text-align: left; border-bottom: 1px solid var(--c-border); font-size: 0.9rem; }
.admin-table th { color: var(--c-text-muted); font-weight: 500; text-transform: uppercase; font-size: 0.8rem; letter-spacing: 0.05em; }
.modal { position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); width: 100%; max-width: 600px; padding: var(--sp-8); border-radius: var(--r-xl); max-height: 90vh; overflow-y: auto; }
</style>
