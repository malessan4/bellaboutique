<template>
  <div>
    <!-- Hero -->
    <section class="hero">
      <div class="hero__bg"></div>
      <div class="hero__particles"></div>
      <div class="container hero__content">
        <h1 class="display-title visible-on-scroll">Belleza que enamora</h1>
        <p class="hero__subtitle visible-on-scroll" style="animation-delay: 0.2s;">Descubrí nuestra nueva colección de lencería y moda femenina premium.</p>
        <div class="hero__actions visible-on-scroll" style="animation-delay: 0.4s;">
          <router-link to="/catalogo" class="btn btn-primary btn-lg">Comprar ahora</router-link>
          <router-link to="/catalogo?category=lenceria" class="btn btn-outline btn-lg">Ver Lencería</router-link>
        </div>
      </div>
      <div class="hero__decor hero__decor-1">♡</div>
      <div class="hero__decor hero__decor-2">✿</div>
    </section>

    <!-- Beneficios -->
    <section class="section benefits">
      <div class="container">
        <div class="benefits__grid">
          <div class="benefit visible-on-scroll">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="1" y="3" width="15" height="13"/><polygon points="16 8 20 8 23 11 23 16 16 16 16 8"/><circle cx="5.5" cy="18.5" r="2.5"/><circle cx="18.5" cy="18.5" r="2.5"/></svg>
            <p>Envío gratis +$30k</p>
          </div>
          <div class="benefit visible-on-scroll" style="animation-delay: 0.1s;">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0110 0v4"/></svg>
            <p>Pagos Seguros</p>
          </div>
          <div class="benefit visible-on-scroll" style="animation-delay: 0.2s;">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>
            <p>Devoluciones</p>
          </div>
          <div class="benefit visible-on-scroll" style="animation-delay: 0.3s;">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>
            <p>Calidad Premium</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Categorías -->
    <section class="section">
      <div class="container">
        <h2 class="section-title visible-on-scroll" style="text-align:center;margin-bottom:var(--sp-12)">Nuestras Categorías</h2>
        <div class="cat-grid">
          <router-link to="/catalogo?category=lenceria" class="cat-card visible-on-scroll">
            <div class="cat-card__bg" style="background-image:url('https://images.unsplash.com/photo-1616086700868-b7c1cd8d3220?w=800&q=80')"></div>
            <div class="cat-card__overlay">
              <h3 class="cat-card__title">Lencería</h3>
              <span class="btn btn-primary">Ver más</span>
            </div>
          </router-link>
          <router-link to="/catalogo?category=ropa-casual" class="cat-card visible-on-scroll" style="animation-delay:0.1s">
            <div class="cat-card__bg" style="background-image:url('https://images.unsplash.com/photo-1515886657613-9f3515b0c78f?w=800&q=80')"></div>
            <div class="cat-card__overlay">
              <h3 class="cat-card__title">Ropa Casual</h3>
              <span class="btn btn-primary">Ver más</span>
            </div>
          </router-link>
          <router-link to="/catalogo?category=pijamas-loungewear" class="cat-card visible-on-scroll" style="animation-delay:0.2s">
            <div class="cat-card__bg" style="background-image:url('https://images.unsplash.com/photo-1601288496920-b6154fe3626a?w=800&q=80')"></div>
            <div class="cat-card__overlay">
              <h3 class="cat-card__title">Pijamas</h3>
              <span class="btn btn-primary">Ver más</span>
            </div>
          </router-link>
          <router-link to="/catalogo?category=accesorios" class="cat-card visible-on-scroll" style="animation-delay:0.3s">
            <div class="cat-card__bg" style="background-image:url('https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=800&q=80')"></div>
            <div class="cat-card__overlay">
              <h3 class="cat-card__title">Accesorios</h3>
              <span class="btn btn-primary">Ver más</span>
            </div>
          </router-link>
        </div>
      </div>
    </section>

    <!-- Destacados -->
    <section class="section">
      <div class="container">
        <h2 class="section-title visible-on-scroll" style="text-align:center;margin-bottom:var(--sp-12)">Productos Destacados</h2>
        <div class="products-grid">
          <template v-if="loading">
            <div v-for="i in 8" :key="i" class="skeleton" style="aspect-ratio:3/4"></div>
          </template>
          <template v-else>
            <ProductCard v-for="p in products" :key="p.id" :product="p" class="visible-on-scroll" />
          </template>
        </div>
      </div>
    </section>

    <!-- Banner -->
    <div class="shipping-banner visible-on-scroll">
      <p>🚚 ¡Envío GRATIS en compras superiores a $30.000!</p>
    </div>

    <!-- Newsletter -->
    <section class="section">
      <div class="container">
        <div class="newsletter glass visible-on-scroll">
          <h2 class="section-title">Unite a Bella Boutique</h2>
          <p style="color:var(--c-text-muted);margin:var(--sp-4) 0 var(--sp-6)">Suscribite y recibí un 10% de descuento en tu primera compra.</p>
          <form class="newsletter__form" @submit.prevent="sub">
            <input type="email" class="form-input" placeholder="Tu correo electrónico" required>
            <button type="submit" class="btn btn-primary">Suscribirme</button>
          </form>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import ProductCard from '@/components/ui/ProductCard.vue'
import api from '@/api' // Assuming a generic api service

const products = ref([])
const loading = ref(true)

async function load() {
  try {
    const res = await api.get('/products/featured')
    products.value = res.data?.data || []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

function sub() { alert('¡Gracias por suscribirte!') }

onMounted(() => {
  load()
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.1 })
  
  document.querySelectorAll('.visible-on-scroll').forEach(el => observer.observe(el))
})
</script>

<style scoped>
.visible-on-scroll { opacity: 0; transform: translateY(24px); transition: opacity 0.6s ease, transform 0.6s ease; }
.visible-on-scroll.visible { opacity: 1; transform: translateY(0); }

.hero { position: relative; min-height: 100vh; display: flex; align-items: center; justify-content: center; text-align: center; overflow: hidden; padding-top: 72px; }
.hero__bg { position: absolute; inset: 0; background: radial-gradient(circle at center, #1a0f12 0%, #0d0a0b 100%); z-index: -2; }
.hero__particles { position: absolute; inset: 0; background-image: radial-gradient(var(--c-rose-dark) 1px, transparent 1px); background-size: 50px 50px; opacity: 0.1; z-index: -1; animation: slideInRight 60s linear infinite; }
.hero__content { position: relative; z-index: 1; display: flex; flex-direction: column; align-items: center; gap: var(--sp-6); }
.hero__subtitle { font-size: 1.25rem; color: var(--c-text-muted); max-width: 600px; margin: 0 auto; }
.hero__actions { display: flex; gap: var(--sp-4); }

.hero__decor { position: absolute; color: var(--c-rose-dark); opacity: 0.2; animation: float 6s ease-in-out infinite; font-size: 4rem; z-index: 0; }
.hero__decor-1 { top: 20%; left: 10%; }
.hero__decor-2 { bottom: 20%; right: 10%; animation-delay: -3s; }

.benefits__grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: var(--sp-6); text-align: center; }
@media (max-width: 768px) { .benefits__grid { grid-template-columns: repeat(2, 1fr); } }
.benefit { display: flex; flex-direction: column; align-items: center; gap: var(--sp-3); padding: var(--sp-6); background: var(--c-card); border-radius: var(--r-lg); border: 1px solid var(--c-border); }
.benefit svg { color: var(--c-rose); }
.benefit p { font-weight: 500; font-size: 0.9rem; }

.cat-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: var(--sp-6); }
@media (max-width: 768px) { .cat-grid { grid-template-columns: 1fr; } }
.cat-card { position: relative; aspect-ratio: 16/9; border-radius: var(--r-lg); overflow: hidden; display: block; }
.cat-card__bg { position: absolute; inset: 0; background-size: cover; background-position: center; transition: transform 0.6s ease; }
.cat-card:hover .cat-card__bg { transform: scale(1.05); }
.cat-card__overlay { position: absolute; inset: 0; background: linear-gradient(0deg, rgba(13,10,11,0.8) 0%, rgba(13,10,11,0.2) 100%); display: flex; flex-direction: column; align-items: center; justify-content: flex-end; padding: var(--sp-8); text-align: center; }
.cat-card__title { font-family: var(--font-display); font-size: 2rem; color: white; margin-bottom: var(--sp-4); transform: translateY(20px); transition: transform 0.4s ease; }
.cat-card .btn { opacity: 0; transform: translateY(20px); transition: all 0.4s ease; }
.cat-card:hover .cat-card__title { transform: translateY(0); }
.cat-card:hover .btn { opacity: 1; transform: translateY(0); }

.shipping-banner { background: linear-gradient(135deg, var(--c-rose), var(--c-rose-dark)); color: white; text-align: center; padding: var(--sp-4); font-weight: 600; letter-spacing: 0.05em; margin: var(--sp-12) 0; }

.newsletter { padding: var(--sp-12); text-align: center; border-radius: var(--r-xl); }
.newsletter__form { display: flex; gap: var(--sp-4); max-width: 500px; margin: 0 auto; }
@media (max-width: 480px) { .newsletter__form { flex-direction: column; } }
</style>
