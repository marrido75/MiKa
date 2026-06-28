<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { productApi, type Product } from '../api/product'
import { orderApi } from '../api/order'
import { useCartStore } from '../stores/cart'

const route = useRoute()
const router = useRouter()
const message = useMessage()
const cart = useCartStore()

const product = ref<Product | null>(null)
const loading = ref(true)
const quantity = ref(1)

const fetchProduct = async () => {
  loading.value = true
  try {
    const res = await productApi.detail(Number(route.params.id))
    product.value = res.data
  } catch {
    message.error('商品不存在')
    router.push('/')
  } finally {
    loading.value = false
  }
}

const addToCart = () => {
  if (!product.value) return
  cart.addItem(product.value, quantity.value)
  message.success('已加入购物车')
}

const buyNow = async () => {
  if (!product.value) return
  try {
    await orderApi.create({
      product_id: product.value.id,
      quantity: quantity.value,
    })
    message.success('下单成功')
    router.push('/order')
  } catch {
    message.error('下单失败，请重试')
  }
}

onMounted(fetchProduct)
</script>

<template>
  <div class="product-detail container">
    <div v-if="loading" class="loading-state">加载中...</div>
    <template v-else-if="product">
      <div class="detail-grid">
        <div class="detail-image">
          <div class="image-placeholder">
            <span class="image-icon">✦</span>
          </div>
        </div>

        <div class="detail-info">
          <span class="detail-tag tag" :class="`tag-${product.category}`">{{ product.category }}</span>
          <h1 class="detail-name">{{ product.name }}</h1>
          <p class="detail-price">¥{{ product.price.toFixed(2) }}</p>
          <p class="detail-stock">库存: {{ product.stock }}</p>
          <p class="detail-desc">{{ product.description }}</p>

          <div class="detail-qty">
            <span class="qty-label">数量</span>
            <div class="qty-controls">
              <button class="qty-btn" :disabled="quantity <= 1" @click="quantity--">−</button>
              <span class="qty-value">{{ quantity }}</span>
              <button class="qty-btn" :disabled="quantity >= product.stock" @click="quantity++">+</button>
            </div>
          </div>

          <div class="detail-actions">
            <button class="btn btn-outline action-btn" @click="addToCart">加入购物车</button>
            <button class="btn btn-accent action-btn" @click="buyNow">立即购买</button>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.product-detail {
  padding: 40px 24px 60px;
}

.detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 48px;
  align-items: start;
}

.detail-image {
  position: sticky;
  top: 88px;
}

.image-placeholder {
  aspect-ratio: 1;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
}

.image-icon {
  font-size: 80px;
  color: rgba(255, 255, 255, 0.8);
}

.detail-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.detail-tag {
  align-self: flex-start;
}

.detail-name {
  font-family: var(--font-serif);
  font-size: 28px;
  font-weight: 700;
  color: var(--color-text);
}

.detail-price {
  font-size: 28px;
  font-weight: 700;
  color: var(--color-accent);
}

.detail-stock {
  font-size: 14px;
  color: var(--color-text-secondary);
}

.detail-desc {
  font-size: 15px;
  color: var(--color-text-secondary);
  line-height: 1.8;
}

.detail-qty {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-top: 8px;
}

.qty-label {
  font-size: 14px;
  color: var(--color-text-secondary);
}

.qty-controls {
  display: flex;
  align-items: center;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-sm);
  overflow: hidden;
}

.qty-btn {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  background: var(--color-bg);
  color: var(--color-text);
  font-size: 18px;
  cursor: pointer;
  transition: background-color 0.15s ease;
}

.qty-btn:hover:not(:disabled) {
  background-color: var(--color-border);
}

.qty-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.qty-value {
  width: 48px;
  text-align: center;
  font-size: 15px;
  font-weight: 500;
}

.detail-actions {
  display: flex;
  gap: 16px;
  margin-top: 16px;
}

.action-btn {
  flex: 1;
}

.loading-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}

@media (max-width: 768px) {
  .detail-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }

  .detail-image {
    position: static;
  }
}
</style>
