import api from './index'

export interface Product {
  id: number
  name: string
  description: string
  price: number
  category: string
  image_url: string
  stock: number
  status: string
  created_at: string
}

export const productApi = {
  list(category?: string) {
    return api.get<Product[]>('/products', { params: { category } })
  },

  detail(id: number) {
    return api.get<Product>(`/products/${id}`)
  },
}
