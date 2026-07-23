<template>
  <div class="admin-layout">
    <AdminSidebar />
    <main class="admin-main">
      <h1 class="display-title" style="font-size:2rem; margin-bottom:var(--sp-8)">Dashboard</h1>

      <div v-if="loading" style="display:flex;gap:var(--sp-4)">
        <div class="skeleton" style="height:120px;flex:1" v-for="i in 3" :key="i"></div>
      </div>
      <div v-else class="stats-grid">
        <div class="stat-card glass">
          <h3>Total Pedidos</h3>
          <p class="stat-val">{{ stats.total_orders }}</p>
        </div>
        <div class="stat-card glass">
          <h3>Ingresos Totales</h3>
          <p class="stat-val" style="color:var(--c-rose)">{{ fmt(stats.total_revenue) }}</p>
        </div>
        <div class="stat-card glass">
          <h3>Pedidos Pendientes</h3>
          <p class="stat-val" style="color:var(--c-gold)">{{ stats.pending_orders }}</p>
        </div>
      </div>

      <h2 class="section-title" style="font-size:1.5rem; margin:var(--sp-12) 0 var(--sp-6)">Últimos Pedidos</h2>
      <div class="table-wrap glass">
        <table class="admin-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Cliente</th>
              <th>Total</th>
              <th>Estado</th>
              <th>Fecha</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="o in recentOrders" :key="o.id">
              <td>#{{ o.id }}</td>
              <td>{{ o.customer_name }}</td>
              <td>{{ fmt(o.total_amount) }}</td>
              <td><span class="badge" :class="'badge-'+o.status">{{ o.status }}</span></td>
              <td>{{ new Date(o.created_at).toLocaleDateString() }}</td>
            </tr>
            <tr v-if="recentOrders.length===0">
              <td colspan="5" style="text-align:center;color:var(--c-text-muted)">No hay pedidos recientes</td>
            </tr>
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

const loading = ref(true)
const stats = ref({ total_orders: 0, total_revenue: 0, pending_orders: 0 })
const recentOrders = ref([])

function fmt(p){return new Intl.NumberFormat('es-AR',{style:'currency',currency:'ARS',maximumFractionDigits:0}).format(p)}

onMounted(async () => {
  try {
    const res = await api.get('/admin/stats')
    stats.value = res.data.stats
    recentOrders.value = res.data.recent_orders || []
  } catch(e) {
    console.error(e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.admin-layout { display: flex; min-height: 100vh; }
.admin-main { flex: 1; padding: var(--sp-8); overflow-y: auto; background: var(--c-bg); }
.stats-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: var(--sp-6); margin-bottom: var(--sp-12); }
.stat-card { padding: var(--sp-6); border-radius: var(--r-lg); }
.stat-card h3 { font-size: 0.9rem; color: var(--c-text-muted); font-weight: 500; text-transform: uppercase; letter-spacing: 0.05em; margin-bottom: var(--sp-2); }
.stat-val { font-family: var(--font-display); font-size: 2.5rem; font-weight: 600; }

.table-wrap { border-radius: var(--r-lg); overflow-x: auto; }
.admin-table { width: 100%; border-collapse: collapse; }
.admin-table th, .admin-table td { padding: var(--sp-4); text-align: left; border-bottom: 1px solid var(--c-border); font-size: 0.9rem; }
.admin-table th { color: var(--c-text-muted); font-weight: 500; text-transform: uppercase; font-size: 0.8rem; letter-spacing: 0.05em; }
.admin-table tr:last-child td { border-bottom: none; }
</style>
