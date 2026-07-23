<template>
  <div class="fail-page">
    <div class="fail-card glass">
      <div class="icon-wrap">
        <svg class="x-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
      </div>
      
      <h1 class="display-title" style="font-size:2rem; margin-bottom:var(--sp-4)">El pago no pudo procesarse</h1>
      <p style="color:var(--c-text-muted); margin-bottom:var(--sp-8)">Hubo un inconveniente al intentar cobrar tu pedido. Podés intentarlo de nuevo o usar otro método de pago.</p>
      
      <div class="actions">
        <router-link to="/checkout" class="btn btn-primary">Intentar de nuevo</router-link>
        <router-link to="/catalogo" class="btn btn-outline">Seguir comprando</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fail-page { min-height: 80vh; display: flex; align-items: center; justify-content: center; padding: var(--sp-6); }
.fail-card { max-width: 500px; width: 100%; text-align: center; padding: var(--sp-12); border-radius: var(--r-xl); animation: fadeUp 0.4s ease; }
.icon-wrap { width: 80px; height: 80px; border-radius: 50%; background: rgba(255,91,91,0.1); color: #ff5b5b; display: flex; align-items: center; justify-content: center; margin: 0 auto var(--sp-6); animation: shake 0.5s ease 0.2s; }
.x-icon { width: 40px; height: 40px; }
.actions { display: flex; gap: var(--sp-4); justify-content: center; }
@keyframes shake { 0%,100%{transform:translateX(0);} 25%{transform:translateX(-8px);} 75%{transform:translateX(8px);} }
</style>
