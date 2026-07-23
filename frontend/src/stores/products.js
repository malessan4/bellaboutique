import { defineStore } from 'pinia'
import { ref } from 'vue'
import { productsApi, categoriesApi } from '@/api'

export const useProductStore = defineStore('products', () => {
  const products = ref([])
  const categories = ref([])
  const featured = ref([])
  const loading = ref(false)
  const total = ref(0)
  const pages = ref(1)

  async function fetchProducts(params = {}) {
    loading.value = true
    try {
      const res = await productsApi.getAll(params)
      products.value = res.data.products || []
      total.value = res.data.total || 0
      pages.value = res.data.pages || 1
    } finally {
      loading.value = false
    }
  }

  async function fetchFeatured() {
    try {
      const res = await productsApi.getFeatured()
      featured.value = res.data || []
    } catch {}
  }

  async function fetchCategories() {
    try {
      const res = await categoriesApi.getAll()
      categories.value = res.data || []
    } catch {}
  }

  return { products, categories, featured, loading, total, pages, fetchProducts, fetchFeatured, fetchCategories }
})
