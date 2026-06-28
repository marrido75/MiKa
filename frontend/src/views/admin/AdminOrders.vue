<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'

interface AdminOrder {
  id: number
  order_no: string
  product_name: string
  quantity: number
  total_price: number
  status: string
  created_at: string
}

const orders = ref<AdminOrder[]>([])
const loading = ref(true)

const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await api.get<AdminOrder[]>('/admin/orders')
    orders.value = res.data
  } catch {
    orders.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchOrders)
</script>

<template>
  <div class="admin-orders">
    <div class="page-header">
      <h1>订单管理</h1>
    </div>

    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="orders.length === 0" class="empty-state">暂无订单</div>
    <table v-else class="data-table">
      <thead>
        <tr>
          <th>订单号</th>
          <th>商品</th>
          <th>数量</th>
          <th>金额</th>
          <th>状态</th>
          <th>时间</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="o in orders" :key="o.id">
          <td class="order-no">{{ o.order_no }}</td>
          <td>{{ o.product_name }}</td>
          <td>{{ o.quantity }}</td>
          <td>¥{{ o.total_price.toFixed(2) }}</td>
          <td>
            <span class="status-tag" :class="o.status">{{ o.status }}</span>
          </td>
          <td>{{ new Date(o.created_at).toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.admin-orders {
  max-width: 1100px;
  margin: 0 auto;
  padding: 40px 24px;
}

.page-header {
  margin-bottom: 32px;
}

.page-header h1 {
  font-size: 24px;
  font-weight: 600;
  color: var(--color-text);
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  background: var(--color-card);
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.06);
}

.data-table th,
.data-table td {
  padding: 12px 16px;
  text-align: left;
  font-size: 14px;
  border-bottom: 1px solid var(--color-border);
}

.data-table th {
  background: var(--color-bg);
  font-weight: 600;
  color: var(--color-text-secondary);
}

.order-no {
  font-family: monospace;
  font-size: 13px;
}

.status-tag {
  display: inline-block;
  padding: 2px 10px;
  border-radius: var(--radius-full);
  font-size: 12px;
  font-weight: 500;
  background: var(--color-bg);
  color: var(--color-text-secondary);
}

.status-tag.pending {
  background: #fef3cd;
  color: #856404;
}

.status-tag.paid {
  background: #d4edda;
  color: #155724;
}

.status-tag.completed {
  background: #cce5ff;
  color: #004085;
}

.status-tag.failed {
  background: #f8d7da;
  color: #721c24;
}

.loading-state,
.empty-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}
</style>
