import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
  timeout: 15000,
  headers: { 'Content-Type': 'application/json' }
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem('bb_token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

api.interceptors.response.use(
  res => res,
  err => {
    if (err.response?.status === 401) {
      localStorage.removeItem('bb_token')
      localStorage.removeItem('bb_user')
      if (window.location.pathname.startsWith('/admin') && !window.location.pathname.includes('/login')) {
        window.location.href = '/admin/login'
      }
    }
    return Promise.reject(err)
  }
)

export default api

export const productsApi = {
  getAll: (params) => api.get('/api/products', { params }),
  getFeatured: () => api.get('/api/products/featured'),
  getBySlug: (slug) => api.get(`/api/products/${slug}`),
  create: (data) => api.post('/api/admin/products', data),
  update: (id, data) => api.put(`/api/admin/products/${id}`, data),
  delete: (id) => api.delete(`/api/admin/products/${id}`),
  uploadImage: (file) => {
    const fd = new FormData()
    fd.append('image', file)
    return api.post('/api/admin/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  }
}

export const categoriesApi = {
  getAll: () => api.get('/api/categories'),
}

export const ordersApi = {
  create: (data) => api.post('/api/orders', data),
  getById: (id) => api.get(`/api/orders/${id}`),
  getAll: (params) => api.get('/api/admin/orders', { params }),
  update: (id, data) => api.put(`/api/admin/orders/${id}`, data),
  getStats: () => api.get('/api/admin/stats'),
}

export const paymentsApi = {
  create: (orderId) => api.post('/api/payments/create', { order_id: orderId }),
}

export const authApi = {
  login: (email, password) => api.post('/api/auth/login', { email, password }),
}
