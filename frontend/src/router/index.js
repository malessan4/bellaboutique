import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', component: () => import('@/views/HomeView.vue'), meta: { title: 'Bella Boutique — Moda Femenina & Lencería' } },
  { path: '/catalogo', component: () => import('@/views/CatalogView.vue'), meta: { title: 'Catálogo — Bella Boutique' } },
  { path: '/producto/:slug', component: () => import('@/views/ProductView.vue') },
  { path: '/carrito', component: () => import('@/views/CartView.vue'), meta: { title: 'Carrito — Bella Boutique' } },
  { path: '/checkout', component: () => import('@/views/CheckoutView.vue'), meta: { title: 'Checkout — Bella Boutique' } },
  { path: '/pago/exitoso', component: () => import('@/views/PaymentSuccess.vue'), meta: { title: '¡Pago Exitoso! — Bella Boutique' } },
  { path: '/pago/fallido', component: () => import('@/views/PaymentFailure.vue'), meta: { title: 'Pago Fallido — Bella Boutique' } },
  { path: '/pago/pendiente', component: () => import('@/views/PaymentPending.vue'), meta: { title: 'Pago Pendiente — Bella Boutique' } },
  { path: '/admin/login', component: () => import('@/views/admin/AdminLogin.vue'), meta: { layout: 'plain', title: 'Admin — Bella Boutique' } },
  { path: '/admin', component: () => import('@/views/admin/AdminDashboard.vue'), meta: { requiresAuth: true, title: 'Dashboard — Admin' } },
  { path: '/admin/productos', component: () => import('@/views/admin/AdminProducts.vue'), meta: { requiresAuth: true, title: 'Productos — Admin' } },
  { path: '/admin/pedidos', component: () => import('@/views/admin/AdminOrders.vue'), meta: { requiresAuth: true, title: 'Pedidos — Admin' } },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, saved) {
    if (saved) return saved
    return { top: 0, behavior: 'smooth' }
  }
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) document.title = to.meta.title
  if (to.meta.requiresAuth && !localStorage.getItem('bb_token')) {
    return next('/admin/login')
  }
  next()
})

export default router
