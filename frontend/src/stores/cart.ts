import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Product } from '../api/product'

export interface CartItem {
  product: Product
  quantity: number
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>(JSON.parse(localStorage.getItem('cart') || '[]'))

  const total = computed(() =>
    items.value.reduce((sum, item) => sum + item.product.price * item.quantity, 0)
  )

  const totalCount = computed(() =>
    items.value.reduce((sum, item) => sum + item.quantity, 0)
  )

  const addItem = (product: Product, quantity = 1) => {
    const existing = items.value.find(i => i.product.id === product.id)
    if (existing) {
      existing.quantity += quantity
    } else {
      items.value.push({ product, quantity })
    }
    save()
  }

  const removeItem = (productId: number) => {
    items.value = items.value.filter(i => i.product.id !== productId)
    save()
  }

  const updateQuantity = (productId: number, quantity: number) => {
    const item = items.value.find(i => i.product.id === productId)
    if (item) {
      item.quantity = quantity
      save()
    }
  }

  const clear = () => {
    items.value = []
    save()
  }

  const save = () => {
    localStorage.setItem('cart', JSON.stringify(items.value))
  }

  return { items, total, totalCount, addItem, removeItem, updateQuantity, clear }
})
