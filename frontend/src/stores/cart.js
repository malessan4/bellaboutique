import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useCartStore = defineStore('cart', () => {
  const items = ref([])
  const isOpen = ref(false)

  try {
    const saved = localStorage.getItem('bb_cart')
    if (saved) items.value = JSON.parse(saved)
  } catch {}

  function persist() { localStorage.setItem('bb_cart', JSON.stringify(items.value)) }

  const count = computed(() => items.value.reduce((s, i) => s + i.quantity, 0))
  const subtotal = computed(() => items.value.reduce((s, i) => s + i.price * i.quantity, 0))
  const shipping = computed(() => subtotal.value >= 30000 ? 0 : 2500)
  const total = computed(() => subtotal.value + shipping.value)

  function addItem(product, size, color, quantity = 1) {
    const idx = items.value.findIndex(i => i.id === product.id && i.size === size && i.color === color)
    if (idx >= 0) {
      items.value[idx].quantity += quantity
    } else {
      items.value.push({
        id: product.id,
        name: product.name,
        price: product.sale_price || product.price,
        image: product.images?.[0] || '',
        size,
        color,
        quantity,
        slug: product.slug,
        category: product.category?.name || ''
      })
    }
    isOpen.value = true
    persist()
  }

  function removeItem(id, size, color) {
    items.value = items.value.filter(i => !(i.id === id && i.size === size && i.color === color))
    persist()
  }

  function updateQuantity(id, size, color, quantity) {
    const item = items.value.find(i => i.id === id && i.size === size && i.color === color)
    if (item) {
      if (quantity <= 0) removeItem(id, size, color)
      else { item.quantity = quantity; persist() }
    }
  }

  function clear() {
    items.value = []
    localStorage.removeItem('bb_cart')
  }

  return { items, isOpen, count, subtotal, shipping, total, addItem, removeItem, updateQuantity, clear }
})
