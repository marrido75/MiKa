import api from './index'

export interface User {
  id: number
  username: string
  email: string
  role: string
}

export interface AuthResponse {
  token: string
  user: User
}

export const authApi = {
  register(data: { username: string; email: string; password: string }) {
    return api.post<AuthResponse>('/auth/register', data)
  },

  login(data: { username: string; password: string }) {
    return api.post<AuthResponse>('/auth/login', data)
  },

  getProfile() {
    return api.get<User>('/auth/profile')
  },
}
