import axios from 'axios'

function normalizeKeys(obj: any): any {
  if (obj === null || obj === undefined) return obj
  if (Array.isArray(obj)) return obj.map(normalizeKeys)
  if (typeof obj !== 'object') return obj

  const result: any = {}
  for (const key of Object.keys(obj)) {
    const lowerKey = key.charAt(0).toLowerCase() + key.slice(1)
    result[lowerKey] = normalizeKeys(obj[key])
  }
  return result
}

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (response) => {
    response.data = normalizeKeys(response.data)
    return response
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.hash = '#/login'
    }
    return Promise.reject(error.response?.data || error)
  }
)

export default api
