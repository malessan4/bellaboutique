<template>
  <div class="pend-page">
    <div class="pend-card glass">
      <div class="icon-wrap">
        <svg class="clock-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
      </div>
      
      <h1 class="display-title" style="font-size:2rem; margin-bottom:var(--sp-4)">Tu pago está siendo procesado</h1>
      <p style="color:var(--c-text-muted); margin-bottom:var(--sp-8)">En breve recibirás la confirmación por email una vez que se acredite el pago.</p>
      
      <router-link to="/" class="btn btn-primary">Volver al inicio</router-link>
    </div>
  </div>
</template>

<style scoped>
.pend-page { min-height: 80vh; display: flex; align-items: center; justify-content: center; padding: var(--sp-6); }
.pend-card { max-width: 500px; width: 100%; text-align: center; padding: var(--sp-12); border-radius: var(--r-xl); animation: fadeUp 0.4s ease; }
.icon-wrap { width: 80px; height: 80px; border-radius: 50%; background: rgba(201,169,110,0.1); color: var(--c-gold); display: flex; align-items: center; justify-content: center; margin: 0 auto var(--sp-6); }
.clock-icon { width: 40px; height: 40px; animation: spin 4s linear infinite; }
</style>
