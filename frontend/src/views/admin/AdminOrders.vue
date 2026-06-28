<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'

interface OrderCard {
  id: number
  content: string
  status: string
}

interface AdminOrder {
  id: number
  order_no: string
  product_name: string
  quantity: number
  total_price: number
  discount: number
  coupon_code: string
  status: string
  created_at: string
  cards: OrderCard[]
}

const orders = ref<AdminOrder[]>([])
const loading = ref(true)
const expandedId = ref<number | null>(null)

const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await api.get<{ orders: AdminOrder[] }>('/admin/orders')
    orders.value = res.data.orders || res.data
  } catch {
    orders.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchOrders)

const toggleRow = (id: number) => {
  expandedId.value = expandedId.value === id ? null : id
}
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
          <th></th>
          <th>订单号</th>
          <th>商品</th>
          <th>数量</th>
          <th>金额</th>
          <th>状态</th>
          <th>时间</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="o in orders" :key="o.id">
          <tr class="order-row" :class="{ expanded: expandedId === o.id }" @click="toggleRow(o.id)">
            <td class="expand-icon">{{ expandedId === o.id ? '▼' : '▶' }}</td>
            <td class="order-no">{{ o.order_no }}</td>
            <td>{{ o.product_name }}</td>
            <td>{{ o.quantity }}</td>
            <td>¥{{ o.total_price.toFixed(2) }}</td>
            <td>
              <span class="status-tag" :class="o.status">{{ o.status }}</span>
            </td>
            <td>{{ new Date(o.created_at).toLocaleString() }}</td>
          </tr>
          <tr v-if="expandedId === o.id" class="detail-row">
            <td colspan="7">
              <div class="order-detail">
                <div class="detail-fields">
                  <div class="detail-item">
                    <span class="detail-label">折扣</span>
                    <span>¥{{ (o.discount || 0).toFixed(2) }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">优惠码</span>
                    <span>{{ o.coupon_code || '无' }}</span>
                  </div>
                </div>
                <div v-if="o.cards && o.cards.length > 0" class="detail-cards">
                  <span class="detail-label">卡密</span>
                  <div class="card-codes">
                    <span v-for="card in o.cards" :key="card.id" class="card-code-tag">{{ card.content }}</span>
                  </div>
                </div>
              </div>
            </td>
          </tr>
        </template>
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

.order-row {
  cursor: pointer;
  transition: background 0.15s;
}

.order-row:hover {
  background: var(--color-bg);
}

.order-row.expanded {
  background: var(--color-bg);
}

.expand-icon {
  width: 30px;
  font-size: 12px;
  color: var(--color-text-secondary);
}

.order-no {
  font-family: monospace;
  font-size: 13px;
}

.detail-row td {
  padding: 0;
  border-bottom: 1px solid var(--color-border);
}

.order-detail {
  padding: 16px 16px 16px 46px;
  background: var(--color-bg);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-fields {
  display: flex;
  gap: 32px;
}

.detail-item {
  display: flex;
  gap: 8px;
  font-size: 14px;
}

.detail-label {
  font-weight: 600;
  color: var(--color-text-secondary);
  font-size: 13px;
}

.detail-cards {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.card-codes {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.card-code-tag {
  font-family: monospace;
  font-size: 12px;
  padding: 3px 10px;
  background: var(--color-card);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-full);
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
