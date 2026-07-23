<template>
  <header class="nav" :class="{scrolled:scrolled}">
    <div class="container nav__inner">
      <router-link to="/" class="nav__logo">
        <img src="/logo.png" alt="Bella Boutique" class="nav__logo-img" @error="logoErr=true" v-if="!logoErr"/>
        <span class="nav__logo-txt" v-else>Bella Boutique</span>
      </router-link>

      <nav class="nav__links hide-mobile">
        <router-link to="/" class="nav__a">Inicio</router-link>
        <router-link to="/catalogo" class="nav__a">Catálogo</router-link>
        <router-link to="/catalogo?category=lenceria" class="nav__a">Lencería</router-link>
        <router-link to="/catalogo?category=ropa-casual" class="nav__a">Ropa</router-link>
        <router-link to="/catalogo?category=pijamas-loungewear" class="nav__a">Pijamas</router-link>
        <router-link to="/catalogo?category=accesorios" class="nav__a">Accesorios</router-link>
      </nav>

      <div class="nav__actions">
        <button class="nav__icon hide-mobile" @click="toggleSearch" aria-label="Buscar">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
        </button>
        <button class="nav__icon nav__cart" @click="cartStore.isOpen=true" aria-label="Carrito">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 01-8 0"/></svg>
          <span class="nav__badge" v-if="cartStore.count>0">{{ cartStore.count }}</span>
        </button>
        <button class="nav__icon hide-desktop" @click="mOpen=true" aria-label="Menú">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/></svg>
        </button>
      </div>
    </div>

    <div class="nav__search" v-if="sOpen">
      <div class="container">
        <input ref="sinput" v-model="sq" class="form-input" placeholder="Buscar productos..." @keyup.enter="doSearch" @keyup.esc="sOpen=false" />
      </div>
    </div>
  </header>

  <div class="overlay" v-if="mOpen" @click="mOpen=false"></div>
  <div class="mdrawer" :class="{open:mOpen}">
    <button class="mdrawer__close" @click="mOpen=false">✕</button>
    <span class="nav__logo-txt" style="font-size:1.5rem;color:var(--c-rose)">Bella Boutique</span>
    <nav class="mdrawer__nav">
      <router-link to="/" @click="mOpen=false">Inicio</router-link>
      <router-link to="/catalogo" @click="mOpen=false">Todo el Catálogo</router-link>
      <div style="height:1px;background:var(--c-border)"></div>
      <small style="color:var(--c-text-subtle);letter-spacing:.1em;text-transform:uppercase">Categorías</small>
      <router-link to="/catalogo?category=lenceria" @click="mOpen=false">Lencería</router-link>
      <router-link to="/catalogo?category=ropa-casual" @click="mOpen=false">Ropa Casual</router-link>
      <router-link to="/catalogo?category=pijamas-loungewear" @click="mOpen=false">Pijamas & Loungewear</router-link>
      <router-link to="/catalogo?category=accesorios" @click="mOpen=false">Accesorios</router-link>
    </nav>
  </div>

  <CartDrawer />
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import CartDrawer from '@/components/ui/CartDrawer.vue'

const cartStore = useCartStore()
const router = useRouter()
const scrolled = ref(false)
const mOpen = ref(false)
const sOpen = ref(false)
const sq = ref('')
const sinput = ref(null)
const logoErr = ref(false)

function onScroll() { scrolled.value = window.scrollY > 50 }
onMounted(() => window.addEventListener('scroll', onScroll, { passive: true }))
onUnmounted(() => window.removeEventListener('scroll', onScroll))

async function toggleSearch() {
  sOpen.value = !sOpen.value
  if (sOpen.value) { await nextTick(); sinput.value?.focus() }
}
function doSearch() {
  if (sq.value.trim()) { router.push(`/catalogo?search=${encodeURIComponent(sq.value)}`); sOpen.value=false; sq.value='' }
}
</script>

<style scoped>
.nav{position:fixed;top:0;left:0;right:0;z-index:200;padding:0.75rem 0;transition:all var(--t-base);}
.nav.scrolled{background:rgba(13,10,11,0.92);backdrop-filter:blur(20px);-webkit-backdrop-filter:blur(20px);border-bottom:1px solid var(--c-border);padding:0.5rem 0;}
.nav__inner{display:flex;align-items:center;justify-content:space-between;gap:var(--sp-6);}
.nav__logo{display:flex;align-items:center;}
.nav__logo-img{height:40px;width:auto;object-fit:contain;}
.nav__logo-txt{font-family:var(--font-display);font-size:1.4rem;font-weight:600;color:var(--c-rose);letter-spacing:-0.02em;}
.nav__links{display:flex;align-items:center;gap:var(--sp-6);}
.nav__a{font-size:0.85rem;color:var(--c-text-muted);transition:color var(--t-fast);}
.nav__a:hover,.nav__a.router-link-active{color:var(--c-text);}
.nav__actions{display:flex;align-items:center;gap:var(--sp-2);}
.nav__icon{position:relative;width:40px;height:40px;display:flex;align-items:center;justify-content:center;border-radius:var(--r-full);color:var(--c-text-muted);transition:all var(--t-fast);}
.nav__icon:hover{color:var(--c-text);background:var(--c-card);}
.nav__badge{position:absolute;top:2px;right:2px;min-width:18px;height:18px;background:var(--c-rose);color:white;border-radius:var(--r-full);font-size:0.65rem;font-weight:700;display:flex;align-items:center;justify-content:center;padding:0 4px;}
.nav__search{padding:var(--sp-4) 0;background:var(--c-bg-2);border-top:1px solid var(--c-border);animation:fadeIn 0.2s ease;}
.mdrawer{position:fixed;top:0;left:0;width:280px;height:100vh;background:var(--c-bg-2);border-right:1px solid var(--c-border);z-index:300;transform:translateX(-100%);transition:transform var(--t-base);padding:var(--sp-6);display:flex;flex-direction:column;gap:var(--sp-6);}
.mdrawer.open{transform:translateX(0);}
.mdrawer__close{align-self:flex-end;color:var(--c-text-muted);font-size:1.2rem;}
.mdrawer__close:hover{color:var(--c-text);}
.mdrawer__nav{display:flex;flex-direction:column;gap:var(--sp-4);}
.mdrawer__nav a{font-size:1rem;color:var(--c-text-muted);padding:var(--sp-2) 0;transition:color var(--t-fast);}
.mdrawer__nav a:hover,.mdrawer__nav a.router-link-active{color:var(--c-rose);}
</style>
