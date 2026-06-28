import { defineStore } from 'pinia'
import { ref } from 'vue'
import { authApi, type User } from '../api/auth'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))

  if (token.value) {
    try {
      user.value = JSON.parse(localStorage.getItem('user') || '')
    } catch {}
  }

  const login = async (email: string, password: string) => {
    const res = await authApi.login({ email, password })
    token.value = res.data.token
    user.value = res.data.user
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
  }

  const register = async (username: string, email: string, password: string) => {
    const res = await authApi.register({ username, email, password })
    token.value = res.data.token
    user.value = res.data.user
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  const isLoggedIn = () => !!token.value
  const isAdmin = () => user.value?.role === 'admin'

  return { user, token, login, register, logout, isLoggedIn, isAdmin }
})
