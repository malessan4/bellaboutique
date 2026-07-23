<template>
  <div class="admin-layout">
    <AdminSidebar />
    <main class="admin-main">
      <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:var(--sp-8)">
        <h1 class="display-title" style="font-size:2rem;">Pedidos</h1>
        <select class="form-select" style="width:auto" v-model="filterStatus" @change="load">
          <option value="">Todos</option>
          <option value="pending">Pendientes</option>
          <option value="paid">Pagados</option>
          <option value="shipped">Enviados</option>
          <option value="cancelled">Cancelados</option>
        </select>
      </div>

      <div class="table-wrap glass">
        <table class="admin-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Cliente</th>
              <th>Email</th>
              <th>Total</th>
              <th>Estado</th>
              <th>Fecha</th>
            </tr>
          </thead>
          <tbody>
            <template v-for="o in orders" :key="o.id">
              <tr @click="expanded===o.id?expanded=null:expanded=o.id" style="cursor:pointer" :style="expanded===o.id?'background:rgba(255,255,255,0.02)':''">
                <td>#{{ o.id }}</td>
                <td>{{ o.customer_name }}</td>
                <td>{{ o.customer_email }}</td>
                <td>{{ fmt(o.total_amount) }}</td>
                <td><span class="badge" :class="'badge-'+o.status">{{ o.status }}</span></td>
                <td>{{ new Date(o.created_at).toLocaleDateString() }}</td>
              </tr>
              <tr v-if="expanded===o.id" class="expand-row">
                <td colspan="6">
                  <div class="order-details">
                    <div style="flex:1">
                      <h4>Ítems</h4>
                      <ul style="list-style:none;padding:0;margin-top:var(--sp-2)">
                        <li v-for="i in o.items" :key="i.id" style="font-size:0.85rem;color:var(--c-text-muted);margin-bottom:4px;">
                          {{ i.quantity }}x {{ i.product?.name || 'Producto eliminado' }} ({{ i.price }}) - {{ i.size }} / {{ i.color }}
                        </li>
                      </ul>
                    </div>
                    <div style="flex:1">
                      <h4>Envío</h4>
                      <p style="font-size:0.85rem;color:var(--c-text-muted)">{{ o.shipping_address }}</p>
                    </div>
                    <div>
                      <h4>Actualizar Estado</h4>
                      <div style="display:flex;gap:var(--sp-2);margin-top:var(--sp-2)">
                        <select class="form-select" v-model="o.newStatus" style="padding:0.25rem 0.5rem">
                          <option value="pending">Pendiente</option>
                          <option value="paid">Pagado</option>
                          <option value="shipped">Enviado</option>
                          <option value="cancelled">Cancelado</option>
                        </select>
                        <button class="btn btn-primary btn-sm" @click="updateStatus(o)">Guardar</button>
                      </div>
                    </div>
                  </div>
                </td>
              </tr>
            </template>
          </tbody>
        </table>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AdminSidebar from '@/components/admin/AdminSidebar.vue'
import api from '@/api'

const orders = ref([])
const filterStatus = ref('')
const expanded = ref(null)

function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p)}

async function load() {
  try {
    const res = await api.get(`/orders${filterStatus.value?'?status='+filterStatus.value:''}`)
    orders.value = res.data.data.map(o => ({...o, newStatus: o.status}))
  } catch(e) { console.error(e) }
}

async function updateStatus(o) {
  try {
    await api.patch(`/orders/${o.id}/status`, { status: o.newStatus })
    o.status = o.newStatus
    alert('Estado actualizado')
  } catch(e) { alert('Error al actualizar') }
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
.expand-row td { border-bottom: 2px solid var(--c-rose); padding: 0; }
.order-details { padding: var(--sp-6); background: rgba(0,0,0,0.2); display: flex; gap: var(--sp-6); }
.order-details h4 { font-size: 0.85rem; text-transform: uppercase; letter-spacing: 0.05em; color: var(--c-rose); }
</style>
