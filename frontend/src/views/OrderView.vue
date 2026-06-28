<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import { orderApi, type Order } from '../api/order'

const message = useMessage()
const orders = ref<Order[]>([])
const loading = ref(true)
const expandedId = ref<number | null>(null)

const fetchOrders = async () => {
  loading.value = true
  try {
    const res = await orderApi.myOrders()
    orders.value = res.data
  } catch {
    message.error('获取订单失败')
  } finally {
    loading.value = false
  }
}

const toggleExpand = (id: number) => {
  expandedId.value = expandedId.value === id ? null : id
}

const copyCard = (content: string) => {
  navigator.clipboard.writeText(content).then(() => {
    message.success('已复制')
  })
}

const statusLabel: Record<string, string> = {
  paid: '已支付',
  completed: '已完成',
  pending: '待支付',
  refunded: '已退款',
}

const statusClass: Record<string, string> = {
  paid: 'status-paid',
  completed: 'status-completed',
  pending: 'status-pending',
  refunded: 'status-refunded',
}

onMounted(fetchOrders)
</script>

<template>
  <div class="order-page container">
    <h1 class="page-title">我的订单</h1>

    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="orders.length === 0" class="empty-state">暂无订单</div>

    <div v-else class="order-list">
      <div
        v-for="order in orders"
        :key="order.id"
        class="order-card"
      >
        <div class="order-header" @click="toggleExpand(order.id)">
          <div class="order-main">
            <span class="order-no">{{ order.order_no }}</span>
            <span class="order-product">{{ order.product_name }}</span>
            <span class="order-qty">x{{ order.quantity }}</span>
          </div>
          <div class="order-right">
            <span class="order-price">¥{{ order.total_price.toFixed(2) }}</span>
            <span class="order-status" :class="statusClass[order.status]">
              {{ statusLabel[order.status] || order.status }}
            </span>
            <span class="expand-icon">{{ expandedId === order.id ? '▲' : '▼' }}</span>
          </div>
        </div>

        <div v-if="expandedId === order.id" class="order-detail">
          <div class="detail-row">
            <span>单价: ¥{{ order.unit_price.toFixed(2) }}</span>
            <span>优惠: ¥{{ order.discount.toFixed(2) }}</span>
            <span>下单时间: {{ new Date(order.created_at).toLocaleString() }}</span>
          </div>

          <div v-if="order.cards && order.cards.length > 0" class="cards-section">
            <h4>卡密信息</h4>
            <div v-for="card in order.cards" :key="card.id" class="card-row">
              <span class="card-content">{{ card.content }}</span>
              <button class="copy-btn" @click.stop="copyCard(card.content)">复制</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.order-page {
  padding: 40px 24px 60px;
}

.page-title {
  font-family: var(--font-serif);
  font-size: 28px;
  font-weight: 700;
  margin-bottom: 32px;
}

.order-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.order-card {
  background: var(--color-card);
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border);
  overflow: hidden;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  cursor: pointer;
  transition: background-color 0.15s ease;
}

.order-header:hover {
  background-color: var(--color-bg);
}

.order-main {
  display: flex;
  align-items: center;
  gap: 16px;
}

.order-no {
  font-size: 13px;
  color: var(--color-text-secondary);
  font-family: monospace;
}

.order-product {
  font-size: 15px;
  font-weight: 500;
  color: var(--color-text);
}

.order-qty {
  font-size: 13px;
  color: var(--color-text-secondary);
}

.order-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.order-price {
  font-size: 16px;
  font-weight: 700;
  color: var(--color-accent);
}

.order-status {
  font-size: 12px;
  padding: 3px 10px;
  border-radius: var(--radius-full);
}

.status-paid {
  background: rgba(168, 230, 207, 0.2);
  color: #2D7A4F;
}

.status-completed {
  background: rgba(168, 230, 207, 0.3);
  color: #1B5E20;
}

.status-pending {
  background: rgba(255, 211, 182, 0.3);
  color: #C67A3C;
}

.status-refunded {
  background: rgba(255, 139, 148, 0.2);
  color: #C44569;
}

.expand-icon {
  font-size: 12px;
  color: var(--color-text-secondary);
}

.order-detail {
  padding: 0 20px 16px;
  border-top: 1px solid var(--color-border);
  padding-top: 16px;
  margin: 0 20px;
  padding-left: 0;
  padding-right: 0;
}

.detail-row {
  display: flex;
  gap: 24px;
  font-size: 13px;
  color: var(--color-text-secondary);
  flex-wrap: wrap;
}

.cards-section {
  margin-top: 16px;
}

.cards-section h4 {
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 8px;
}

.card-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  background: var(--color-bg);
  border-radius: var(--radius-sm);
  margin-bottom: 6px;
}

.card-content {
  font-size: 13px;
  font-family: monospace;
  color: var(--color-text);
}

.copy-btn {
  padding: 4px 10px;
  font-size: 12px;
  color: var(--color-primary-pressed);
  background: transparent;
  border: 1px solid var(--color-primary);
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.2s ease;
}

.copy-btn:hover {
  background-color: var(--color-primary);
  color: #fff;
}

.loading-state,
.empty-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}
</style>
