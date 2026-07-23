<template>
  <aside class="asidebar">
    <div class="asidebar__logo">Bella Boutique</div>
    <p class="asidebar__role">Panel Admin</p>
    <nav class="asidebar__nav">
      <router-link to="/admin" exact-active-class="active" class="asidebar__link">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/></svg>
        Dashboard
      </router-link>
      <router-link to="/admin/productos" class="asidebar__link">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 7H4a2 2 0 00-2 2v10a2 2 0 002 2h16a2 2 0 002-2V9a2 2 0 00-2-2z"/><path d="M16 21V5a2 2 0 00-2-2h-4a2 2 0 00-2 2v16"/></svg>
        Productos
      </router-link>
      <router-link to="/admin/pedidos" class="asidebar__link">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7"/></svg>
        Pedidos
      </router-link>
    </nav>
    <button class="asidebar__logout" @click="doLogout">Cerrar sesión</button>
  </aside>
</template>

<script setup>
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'
const auth = useAuthStore()
const router = useRouter()
function doLogout() { auth.logout(); router.push('/admin/login') }
</script>

<style scoped>
.asidebar{width:240px;min-height:100vh;background:var(--c-bg-2);border-right:1px solid var(--c-border);display:flex;flex-direction:column;padding:var(--sp-6);gap:var(--sp-4);flex-shrink:0;}
.asidebar__logo{font-family:var(--font-display);font-size:1.2rem;font-weight:600;color:var(--c-rose);}
.asidebar__role{font-size:0.75rem;color:var(--c-text-subtle);text-transform:uppercase;letter-spacing:0.08em;margin-top:-var(--sp-2);}
.asidebar__nav{display:flex;flex-direction:column;gap:var(--sp-2);margin-top:var(--sp-4);flex:1;}
.asidebar__link{display:flex;align-items:center;gap:var(--sp-3);padding:var(--sp-3) var(--sp-4);border-radius:var(--r-md);color:var(--c-text-muted);font-size:0.9rem;transition:all var(--t-fast);}
.asidebar__link:hover,.asidebar__link.active{background:rgba(201,132,122,0.1);color:var(--c-rose);border-left:2px solid var(--c-rose);}
.asidebar__logout{padding:var(--sp-3) var(--sp-4);border-radius:var(--r-md);color:var(--c-text-subtle);font-size:0.85rem;text-align:left;transition:all var(--t-fast);}
.asidebar__logout:hover{color:#ff5b5b;background:rgba(255,91,91,0.1);}
</style>
