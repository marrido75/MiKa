<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'

interface User {
  id: number
  username: string
  email: string
  role: string
  created_at: string
}

const users = ref<User[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await api.get('/admin/users')
    users.value = res.data
  } catch {
    users.value = []
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="admin-users">
    <h1>用户管理</h1>
    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="users.length === 0" class="empty-state">暂无用户</div>
    <table v-else class="data-table">
      <thead>
        <tr><th>ID</th><th>用户名</th><th>邮箱</th><th>角色</th><th>注册时间</th></tr>
      </thead>
      <tbody>
        <tr v-for="u in users" :key="u.id">
          <td>{{ u.id }}</td>
          <td>{{ u.username }}</td>
          <td>{{ u.email }}</td>
          <td><span class="role-tag" :class="u.role">{{ u.role }}</span></td>
          <td>{{ new Date(u.created_at).toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.admin-users { max-width: 1100px; margin: 0 auto; padding: 40px 24px; }
.admin-users h1 { font-size: 24px; font-weight: 600; margin-bottom: 32px; }
.data-table { width: 100%; border-collapse: collapse; background: var(--color-card); border-radius: var(--radius-lg); overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.06); }
.data-table th, .data-table td { padding: 12px 16px; text-align: left; font-size: 14px; border-bottom: 1px solid var(--color-border); }
.data-table th { background: var(--color-bg); font-weight: 600; color: var(--color-text-secondary); }
.role-tag { padding: 2px 10px; border-radius: var(--radius-full); font-size: 12px; }
.role-tag.admin { background: #cce5ff; color: #004085; }
.role-tag.user { background: var(--color-bg); color: var(--color-text-secondary); }
.loading-state, .empty-state { text-align: center; padding: 60px 0; color: var(--color-text-secondary); }
</style>
