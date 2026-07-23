<template>
  <div class="login-wrapper">
    <div class="login-card glass">
      <h1 class="login-logo">Bella Boutique</h1>
      <p class="login-sub">Panel Administrativo</p>
      
      <form @submit.prevent="login" class="login-form">
        <div class="form-group">
          <label class="form-label">Email</label>
          <input type="email" class="form-input" v-model="email" required>
        </div>
        <div class="form-group">
          <label class="form-label">Contraseña</label>
          <div style="position:relative">
            <input :type="showPass?'text':'password'" class="form-input" v-model="password" required>
            <button type="button" class="pass-toggle" @click="showPass=!showPass">{{ showPass?'Ocultar':'Mostrar' }}</button>
          </div>
        </div>
        
        <p class="form-error" v-if="error">{{ error }}</p>
        
        <button type="submit" class="btn btn-primary btn-full" :disabled="loading" style="margin-top:var(--sp-4)">
          {{ loading ? 'Iniciando...' : 'Iniciar sesión' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()
const email = ref('')
const password = ref('')
const showPass = ref(false)
const loading = ref(false)
const error = ref('')

async function login() {
  loading.value = true
  error.value = ''
  try {
    const success = await auth.login(email.value, password.value)
    if (success) {
      router.push('/admin')
    } else {
      error.value = 'Credenciales incorrectas'
    }
  } catch (e) {
    if (e.response && e.response.status === 401) {
      error.value = 'Credenciales incorrectas (Error 401)'
    } else {
      error.value = 'Error al conectar con el servidor'
    }
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-wrapper { min-height: 100vh; display: flex; align-items: center; justify-content: center; background: radial-gradient(circle at center, var(--c-bg-2) 0%, var(--c-bg) 100%); padding: var(--sp-6); }
.login-card { width: 100%; max-width: 400px; padding: var(--sp-12); border-radius: var(--r-xl); text-align: center; }
.login-logo { font-family: var(--font-display); font-size: 2.5rem; color: var(--c-rose); margin-bottom: var(--sp-1); }
.login-sub { font-size: 0.85rem; text-transform: uppercase; letter-spacing: 0.1em; color: var(--c-text-muted); margin-bottom: var(--sp-8); }
.login-form { text-align: left; }
.pass-toggle { position: absolute; right: 1rem; top: 50%; transform: translateY(-50%); font-size: 0.8rem; color: var(--c-text-muted); }
.pass-toggle:hover { color: var(--c-text); }
</style>
