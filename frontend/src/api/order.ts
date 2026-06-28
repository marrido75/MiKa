import api from './index'

export interface Card {
  id: number
  content: string
  status: string
}

export interface Order {
  id: number
  order_no: string
  product_id: number
  product_name: string
  quantity: number
  unit_price: number
  total_price: number
  coupon_code: string
  discount: number
  status: string
  cards: Card[]
  created_at: string
}

export const orderApi = {
  create(data: { product_id: number; quantity: number; coupon_code?: string }) {
    return api.post<Order>('/orders', data)
  },

  detail(id: number) {
    return api.get<Order>(`/orders/${id}`)
  },

  myOrders() {
    return api.get<Order[]>('/orders/my')
  },
}
