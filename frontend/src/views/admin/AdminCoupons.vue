<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'

interface Coupon {
  id: number
  code: string
  discount_type: string
  discount_value: number
  min_amount: number
  usage_limit: number
  used_count: number
  status: string
  expires_at: string
}

const coupons = ref<Coupon[]>([])
const loading = ref(true)
const showForm = ref(false)
const form = ref({
  code: '',
  discount_type: 'percent',
  discount_value: 10,
  min_amount: 0,
  usage_limit: 100,
  expires_at: '',
})

const fetchCoupons = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/coupons')
    coupons.value = res.data
  } catch {
    coupons.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchCoupons)

const createCoupon = async () => {
  try {
    await api.post('/admin/coupons', form.value)
    showForm.value = false
    fetchCoupons()
    form.value = { code: '', discount_type: 'percent', discount_value: 10, min_amount: 0, usage_limit: 100, expires_at: '' }
  } catch {
    alert('创建失败')
  }
}

const deleteCoupon = async (id: number) => {
  if (!confirm('确认删除？')) return
  try {
    await api.delete(`/admin/coupons/${id}`)
    fetchCoupons()
  } catch {
    alert('删除失败')
  }
}
</script>

<template>
  <div class="admin-coupons">
    <div class="page-header">
      <h1>优惠券管理</h1>
      <button class="btn-primary" @click="showForm = true">创建优惠券</button>
    </div>

    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="coupons.length === 0" class="empty-state">暂无优惠券</div>
    <table v-else class="data-table">
      <thead>
        <tr><th>券码</th><th>类型</th><th>面值</th><th>最低消费</th><th>已用/总量</th><th>状态</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="c in coupons" :key="c.id">
          <td class="code">{{ c.code }}</td>
          <td>{{ c.discount_type === 'percent' ? '折扣' : '满减' }}</td>
          <td>{{ c.discount_type === 'percent' ? c.discount_value + '%' : '¥' + c.discount_value }}</td>
          <td>¥{{ c.min_amount }}</td>
          <td>{{ c.used_count }} / {{ c.usage_limit }}</td>
          <td><span class="status-tag" :class="c.status">{{ c.status }}</span></td>
          <td><button class="btn-danger" @click="deleteCoupon(c.id)">删除</button></td>
        </tr>
      </tbody>
    </table>

    <div v-if="showForm" class="modal-overlay" @click.self="showForm = false">
      <div class="modal-card">
        <h2>创建优惠券</h2>
        <form @submit.prevent="createCoupon" class="admin-form">
          <label><span>券码</span><input v-model="form.code" required /></label>
          <label><span>类型</span><select v-model="form.discount_type"><option value="percent">折扣</option><option value="fixed">满减</option></select></label>
          <label><span>面值</span><input v-model.number="form.discount_value" type="number" min="0" required /></label>
          <label><span>最低消费</span><input v-model.number="form.min_amount" type="number" min="0" /></label>
          <label><span>使用上限</span><input v-model.number="form.usage_limit" type="number" min="0" /></label>
          <label><span>过期时间</span><input v-model="form.expires_at" type="datetime-local" required /></label>
          <div class="form-actions">
            <button type="button" class="btn-secondary" @click="showForm = false">取消</button>
            <button type="submit" class="btn-primary">创建</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-coupons { max-width: 1100px; margin: 0 auto; padding: 40px 24px; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 32px; }
.page-header h1 { font-size: 24px; font-weight: 600; }
.btn-primary { padding: 8px 20px; background: var(--color-primary); color: #fff; border: none; border-radius: var(--radius-full); cursor: pointer; font-size: 14px; }
.btn-primary:hover { background: var(--color-primary-pressed); }
.btn-secondary { padding: 8px 20px; background: var(--color-card); color: var(--color-text-secondary); border: 1.5px solid var(--color-border); border-radius: var(--radius-full); cursor: pointer; font-size: 14px; }
.btn-danger { background: none; border: none; color: #e74c3c; cursor: pointer; font-size: 14px; }
.code { font-family: monospace; font-weight: 600; }
.data-table { width: 100%; border-collapse: collapse; background: var(--color-card); border-radius: var(--radius-lg); overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.06); }
.data-table th, .data-table td { padding: 12px 16px; text-align: left; font-size: 14px; border-bottom: 1px solid var(--color-border); }
.data-table th { background: var(--color-bg); font-weight: 600; color: var(--color-text-secondary); }
.status-tag { padding: 2px 10px; border-radius: var(--radius-full); font-size: 12px; }
.status-tag.active { background: #d4edda; color: #155724; }
.status-tag.inactive { background: #f8d7da; color: #721c24; }
.modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.4); display: flex; align-items: center; justify-content: center; z-index: 100; }
.modal-card { background: var(--color-card); border-radius: var(--radius-lg); padding: 32px; width: 480px; max-width: 90vw; }
.modal-card h2 { margin-bottom: 24px; font-size: 20px; font-weight: 600; }
.admin-form { display: flex; flex-direction: column; gap: 16px; }
.admin-form label { display: flex; flex-direction: column; gap: 6px; }
.admin-form label span { font-size: 13px; font-weight: 500; color: var(--color-text-secondary); }
.admin-form input, .admin-form select { padding: 10px 12px; border: 1.5px solid var(--color-border); border-radius: var(--radius-md); font-size: 14px; background: var(--color-bg); }
.form-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 8px; }
.loading-state, .empty-state { text-align: center; padding: 60px 0; color: var(--color-text-secondary); }
</style>
