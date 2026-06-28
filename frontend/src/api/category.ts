import api from './index'

export interface Category {
  id: number
  name: string
  label: string
  sort: number
  status?: string
}

export const categoryApi = {
  list() {
    return api.get<Category[]>('/categories')
  },

  adminList() {
    return api.get<Category[]>('/admin/categories')
  },

  create(data: { name: string; label: string; sort: number }) {
    return api.post<Category>('/admin/categories', data)
  },

  update(id: number, data: { name: string; label: string; sort: number }) {
    return api.put<Category>(`/admin/categories/${id}`, data)
  },

  delete(id: number) {
    return api.delete(`/admin/categories/${id}`)
  },
}
