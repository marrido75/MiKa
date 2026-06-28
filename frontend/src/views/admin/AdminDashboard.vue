<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'

interface Stats {
  totalOrders: number
  totalRevenue: number
  totalProducts: number
  totalUsers: number
}

const stats = ref<Stats>({ totalOrders: 0, totalRevenue: 0, totalProducts: 0, totalUsers: 0 })
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await api.get('/admin/stats')
    stats.value = res.data
  } catch {
    // fallback
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="admin-dashboard">
    <h1>控制台</h1>
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">📦</div>
        <div class="stat-info">
          <span class="stat-value">{{ stats.totalOrders }}</span>
          <span class="stat-label">总订单</span>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">💰</div>
        <div class="stat-info">
          <span class="stat-value">¥{{ stats.totalRevenue.toFixed(2) }}</span>
          <span class="stat-label">总收入</span>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">🛍️</div>
        <div class="stat-info">
          <span class="stat-value">{{ stats.totalProducts }}</span>
          <span class="stat-label">商品数</span>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">👥</div>
        <div class="stat-info">
          <span class="stat-value">{{ stats.totalUsers }}</span>
          <span class="stat-label">用户数</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-dashboard { max-width: 1100px; margin: 0 auto; padding: 40px 24px; }
.admin-dashboard h1 { font-size: 24px; font-weight: 600; margin-bottom: 32px; }
.stats-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(240px, 1fr)); gap: 20px; }
.stat-card { background: var(--color-card); border-radius: var(--radius-lg); padding: 24px; display: flex; align-items: center; gap: 16px; box-shadow: 0 1px 3px rgba(0,0,0,0.06); }
.stat-icon { font-size: 32px; }
.stat-info { display: flex; flex-direction: column; }
.stat-value { font-size: 24px; font-weight: 700; color: var(--color-text); }
.stat-label { font-size: 13px; color: var(--color-text-secondary); }
.loading { text-align: center; padding: 60px 0; color: var(--color-text-secondary); }
</style>
