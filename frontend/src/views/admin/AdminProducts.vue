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
              <th @click="sortBy('name')" class="sortable">Nombre <span v-if="sortKey==='name'">{{ sortOrder===1 ? '↑' : '↓' }}</span></th>
              <th @click="sortBy('category')" class="sortable">Categoría <span v-if="sortKey==='category'">{{ sortOrder===1 ? '↑' : '↓' }}</span></th>
              <th @click="sortBy('price')" class="sortable">Precio <span v-if="sortKey==='price'">{{ sortOrder===1 ? '↑' : '↓' }}</span></th>
              <th @click="sortBy('stock')" class="sortable">Stock <span v-if="sortKey==='stock'">{{ sortOrder===1 ? '↑' : '↓' }}</span></th>
              <th @click="sortBy('featured')" class="sortable">Destacado <span v-if="sortKey==='featured'">{{ sortOrder===1 ? '↑' : '↓' }}</span></th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in sortedProducts" :key="p.ID">
              <td>{{ p.name }}</td>
              <td>{{ categories.find(c => c.ID === p.category_id)?.name || '-' }}</td>
              <td>
                <span v-if="p.sale_price" style="text-decoration:line-through;color:var(--c-text-muted);font-size:0.8rem;display:block">{{ fmt(p.price) }}</span>
                <span>{{ p.sale_price ? fmt(p.sale_price) : fmt(p.price) }}</span>
              </td>
              <td>{{ p.stock }}</td>
              <td>{{ p.featured ? 'Sí' : 'No' }}</td>
              <td>
                <button class="btn-ghost" style="padding:0.25rem" @click="openModal(p)">Editar</button>
                <button class="btn-ghost" style="padding:0.25rem;color:#ff5b5b" @click="del(p.ID)">Eliminar</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>

    <!-- Modal -->
    <div class="overlay" v-if="modalOpen">
      <div class="modal glass">
        <h2 style="margin-bottom:var(--sp-6)">{{ form.ID ? 'Editar Producto' : 'Nuevo Producto' }}</h2>
        <form @submit.prevent="save">
          <div class="form-group"><label class="form-label">Nombre</label><input type="text" class="form-input" v-model="form.name" required></div>
          <div class="form-group"><label class="form-label">Descripción</label><textarea class="form-input" v-model="form.description" required></textarea></div>
          <div style="display:flex;gap:var(--sp-4)">
            <div class="form-group" style="flex:1"><label class="form-label">Precio</label><input type="number" class="form-input" v-model.number="form.price" required></div>
            <div class="form-group" style="flex:1"><label class="form-label">Precio Oferta</label><input type="number" class="form-input" v-model.number="form.sale_price"></div>
          </div>
          <div class="form-group">
            <label class="form-label">Categoría</label>
            <select class="form-select" v-model.number="form.category_id" required>
              <option v-for="c in categories" :key="c.ID" :value="c.ID">{{ c.name }}</option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">Imágenes</label>
            <div class="image-upload-wrap">
              <input type="file" multiple accept="image/*" @change="handleImageUpload" class="form-input">
              <div class="uploaded-images" v-if="form.images && form.images.length">
                <div v-for="(img, idx) in form.images" :key="idx" class="uploaded-image">
                  <img :src="img" alt="Uploaded" />
                  <button type="button" @click="removeImage(idx)" class="remove-img-btn">X</button>
                </div>
              </div>
            </div>
          </div>
          <div style="display:flex;gap:var(--sp-4)">
            <div class="form-group" style="flex:1"><label class="form-label">Talles (coma)</label><input type="text" class="form-input" v-model="form.sizesStr"></div>
            <div class="form-group" style="flex:1"><label class="form-label">Colores (coma)</label><input type="text" class="form-input" v-model="form.colorsStr"></div>
          </div>
          <div class="form-group" v-if="form.colorsStr && form.images?.length" style="background:rgba(255,255,255,0.02);padding:1rem;border-radius:var(--r-md);margin-top:0.5rem">
            <label class="form-label" style="margin-bottom:0.5rem;display:block">Asignar imágenes a colores</label>
            <div v-for="c in form.colorsStr.split(',').map(s=>s.trim()).filter(Boolean)" :key="c" style="display:flex;align-items:center;gap:1rem;margin-bottom:0.5rem">
              <span style="min-width:80px;font-size:0.9rem">{{ c }}</span>
              <select class="form-select" v-model="form.color_images[c]" style="flex:1;padding:0.5rem;font-size:0.8rem">
                <option value="">(Usar la galería general)</option>
                <option v-for="img in form.images" :key="img" :value="img">Imagen {{ form.images.indexOf(img) + 1 }}</option>
              </select>
              <div style="width:30px;height:40px;background:var(--c-bg-2);border-radius:4px;overflow:hidden;flex-shrink:0">
                <img v-if="form.color_images[c]" :src="form.color_images[c]" style="width:100%;height:100%;object-fit:cover" />
              </div>
            </div>
          </div>
          <div class="form-group"><label class="form-label">Stock</label><input type="number" class="form-input" v-model.number="form.stock" required></div>
          <div class="form-group" style="flex-direction:row;align-items:center;">
            <input type="checkbox" v-model="form.featured" id="feat"><label for="feat">Destacado</label>
          </div>
          
          <div style="display:flex;justify-content:flex-end;gap:var(--sp-4);margin-top:var(--sp-6)">
            <button type="button" class="btn btn-ghost" @click="modalOpen=false">Cancelar</button>
            <button type="submit" class="btn btn-primary">Guardar</button>
          </div>
        </form>
      </div>
    </div>
    <ConfirmModal v-model="confirmOpen" title="Eliminar Producto" message="¿Estás seguro de que deseas eliminar este producto? Esta acción no se puede deshacer." confirm-text="Eliminar" confirm-class="btn-primary" @confirm="doDel" />
  </div>
</template>

<script setup>
import { ref, onMounted, inject, computed } from 'vue'
import AdminSidebar from '@/components/admin/AdminSidebar.vue'
import ConfirmModal from '@/components/ui/ConfirmModal.vue'
import { productsApi, categoriesApi } from '@/api'

const products = ref([])
const categories = ref([])
const modalOpen = ref(false)
const confirmOpen = ref(false)
const itemToDelete = ref(null)
const form = ref({})
const toast = inject('toast', null)

const sortKey = ref('name')
const sortOrder = ref(1)

function sortBy(key) {
  if (sortKey.value === key) {
    sortOrder.value *= -1
  } else {
    sortKey.value = key
    sortOrder.value = 1
  }
}

const sortedProducts = computed(() => {
  return [...products.value].sort((a, b) => {
    let valA = a[sortKey.value]
    let valB = b[sortKey.value]
    
    if (sortKey.value === 'category') {
      valA = categories.value.find(c => c.ID === a.category_id)?.name || ''
      valB = categories.value.find(c => c.ID === b.category_id)?.name || ''
    }
    if (sortKey.value === 'price') {
      valA = a.sale_price ? a.sale_price : a.price
      valB = b.sale_price ? b.sale_price : b.price
    }
    
    if (valA < valB) return -1 * sortOrder.value
    if (valA > valB) return 1 * sortOrder.value
    return 0
  })
})

function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p||0)}

async function load() {
  try {
    const res = await productsApi.getAll({ limit: 100 })
    products.value = res.data.products || []
    const cRes = await categoriesApi.getAll()
    categories.value = cRes.data
  } catch(e) { console.error(e) }
}

function openModal(p) {
  if (p) {
    form.value = { ...p, sizesStr: p.sizes?.join(',')||'', colorsStr: p.colors?.join(',')||'', images: [...(p.images || [])], color_images: { ...(p.color_images || {}) } }
  } else {
    form.value = { name:'', description:'', price:0, stock:0, featured:false, images: [], sizesStr:'', colorsStr:'', color_images: {} }
  }
  modalOpen.value = true
}

async function handleImageUpload(e) {
  const files = e.target.files
  if (!files.length) return
  
  for (let i = 0; i < files.length; i++) {
    try {
      const res = await productsApi.uploadImage(files[i])
      if (res.data && res.data.url) {
        if (!form.value.images) form.value.images = []
        form.value.images.push(res.data.url)
      }
    } catch (error) {
      toast?.error('Error al subir imagen')
    }
  }
  e.target.value = ''
}

function removeImage(idx) {
  form.value.images.splice(idx, 1)
}

async function save() {
  try {
    // Limpiar color_images vacíos
    const cleanColorImages = {}
    if (form.value.color_images) {
      for (const [k, v] of Object.entries(form.value.color_images)) {
        if (v) cleanColorImages[k] = v
      }
    }

    const payload = { ...form.value, 
      sizes: form.value.sizesStr?form.value.sizesStr.split(',').map(s=>s.trim()).filter(Boolean):[],
      colors: form.value.colorsStr?form.value.colorsStr.split(',').map(s=>s.trim()).filter(Boolean):[],
      color_images: cleanColorImages
    }
    // Si form.sale_price es string vacio, setear a null
    if (payload.sale_price === '') payload.sale_price = null

    if (form.value.ID) {
      await productsApi.update(form.value.ID, payload)
      toast?.success('Producto actualizado')
    } else {
      await productsApi.create(payload)
      toast?.success('Producto creado')
    }
    modalOpen.value = false
    load()
  } catch(e) { toast?.error('Error al guardar el producto') }
}

function del(id) {
  itemToDelete.value = id
  confirmOpen.value = true
}

async function doDel() {
  if (itemToDelete.value) {
    try {
      await productsApi.delete(itemToDelete.value)
      toast?.success('Producto eliminado')
      load()
    } catch (e) {
      toast?.error('Error al eliminar producto')
    }
    itemToDelete.value = null
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
.admin-table th { color: var(--c-text-muted); font-weight: 500; text-transform: uppercase; font-size: 0.8rem; letter-spacing: 0.05em; transition: color var(--t-fast); }
.admin-table th.sortable { cursor: pointer; user-select: none; }
.admin-table th.sortable:hover { color: var(--c-rose); }
.modal { position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); width: 100%; max-width: 600px; padding: var(--sp-8); border-radius: var(--r-xl); max-height: 90vh; overflow-y: auto; }
.image-upload-wrap { display: flex; flex-direction: column; gap: var(--sp-3); }
.uploaded-images { display: flex; gap: var(--sp-3); flex-wrap: wrap; }
.uploaded-image { position: relative; width: 80px; height: 80px; border-radius: var(--r-md); overflow: hidden; border: 1px solid var(--c-border); }
.uploaded-image img { width: 100%; height: 100%; object-fit: cover; }
.remove-img-btn { position: absolute; top: 0; right: 0; background: rgba(0,0,0,0.6); color: white; border: none; width: 24px; height: 24px; display: flex; align-items: center; justify-content: center; cursor: pointer; font-size: 10px; border-bottom-left-radius: var(--r-sm); }
.remove-img-btn:hover { background: red; }
</style>
