import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('bb_token') || '')
  const user = ref(JSON.parse(localStorage.getItem('bb_user') || 'null'))
  const isAuthenticated = computed(() => !!token.value)

  async function login(email, password) {
    const res = await authApi.login(email, password)
    token.value = res.data.token
    user.value = res.data.user
    localStorage.setItem('bb_token', token.value)
    localStorage.setItem('bb_user', JSON.stringify(user.value))
    return true
  }

  function logout() {
    token.value = ''
    user.value = null
    localStorage.removeItem('bb_token')
    localStorage.removeItem('bb_user')
  }

  return { token, user, isAuthenticated, login, logout }
})
