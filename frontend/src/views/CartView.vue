<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useCartStore } from '../stores/cart'
import { orderApi } from '../api/order'
import { couponApi } from '../api/coupon'
import CartItem from '../components/CartItem.vue'

const router = useRouter()
const message = useMessage()
const cart = useCartStore()

const couponCode = ref('')
const discount = ref(0)
const couponMessage = ref('')
const checkingOut = ref(false)

const finalPrice = computed(() => {
  const total = cart.total
  return Math.max(0, total - discount.value)
})

const verifyCoupon = async () => {
  if (!couponCode.value.trim()) {
    message.warning('请输入优惠码')
    return
  }
  try {
    const res = await couponApi.verify(couponCode.value, cart.total)
    if (res.data.valid) {
      discount.value = res.data.discount
      couponMessage.value = res.data.message
      message.success(res.data.message)
    } else {
      discount.value = 0
      couponMessage.value = ''
      message.error(res.data.message)
    }
  } catch {
    message.error('验证优惠码失败')
  }
}

const checkout = async () => {
  if (cart.items.length === 0) {
    message.warning('购物车为空')
    return
  }

  checkingOut.value = true
  try {
    for (const item of cart.items) {
      await orderApi.create({
        product_id: item.product.id,
        quantity: item.quantity,
        coupon_code: couponCode.value || undefined,
      })
    }
    cart.clear()
    message.success('下单成功')
    router.push('/order')
  } catch {
    message.error('下单失败，请重试')
  } finally {
    checkingOut.value = false
  }
}
</script>

<template>
  <div class="cart-page container">
    <h1 class="page-title">购物车</h1>

    <div v-if="cart.items.length === 0" class="empty-state">购物车是空的</div>

    <div v-else class="cart-layout">
      <div class="cart-list">
        <CartItem
          v-for="item in cart.items"
          :key="item.product.id"
          :item="item"
          @remove="cart.removeItem"
          @update="cart.updateQuantity"
        />
      </div>

      <div class="cart-summary">
        <h3 class="summary-title">订单摘要</h3>

        <div class="coupon-row">
          <input
            v-model="couponCode"
            class="coupon-input"
            placeholder="输入优惠码"
          />
          <button class="btn btn-outline btn-sm" @click="verifyCoupon">验证</button>
        </div>
        <p v-if="couponMessage" class="coupon-msg">{{ couponMessage }}</p>

        <div class="summary-line">
          <span>商品总计</span>
          <span>¥{{ cart.total.toFixed(2) }}</span>
        </div>
        <div v-if="discount > 0" class="summary-line discount">
          <span>优惠</span>
          <span>-¥{{ discount.toFixed(2) }}</span>
        </div>
        <div class="summary-total">
          <span>实付</span>
          <span class="total-price">¥{{ finalPrice.toFixed(2) }}</span>
        </div>

        <button
          class="btn btn-accent checkout-btn"
          :disabled="checkingOut"
          @click="checkout"
        >
          {{ checkingOut ? '处理中...' : '结算' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.cart-page {
  padding: 40px 24px 60px;
}

.page-title {
  font-family: var(--font-serif);
  font-size: 28px;
  font-weight: 700;
  margin-bottom: 32px;
}

.cart-layout {
  display: grid;
  grid-template-columns: 1fr 340px;
  gap: 32px;
  align-items: start;
}

.cart-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.cart-summary {
  position: sticky;
  top: 88px;
  padding: 24px;
  background: var(--color-card);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.summary-title {
  font-size: 16px;
  font-weight: 600;
}

.coupon-row {
  display: flex;
  gap: 8px;
}

.coupon-input {
  flex: 1;
  padding: 8px 12px;
  font-size: 13px;
  border: 1.5px solid var(--color-border);
  border-radius: var(--radius-sm);
  background: var(--color-bg);
  color: var(--color-text);
  outline: none;
}

.coupon-input:focus {
  border-color: var(--color-primary);
}

.coupon-msg {
  font-size: 12px;
  color: var(--color-primary-pressed);
}

.summary-line {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
  color: var(--color-text-secondary);
}

.summary-line.discount {
  color: var(--color-primary-pressed);
}

.summary-total {
  display: flex;
  justify-content: space-between;
  font-size: 16px;
  font-weight: 600;
  padding-top: 12px;
  border-top: 1px solid var(--color-border);
}

.total-price {
  color: var(--color-accent);
  font-size: 20px;
}

.checkout-btn {
  width: 100%;
  margin-top: 8px;
}

.empty-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}

.btn-sm {
  padding: 8px 16px;
  font-size: 13px;
}

@media (max-width: 768px) {
  .cart-layout {
    grid-template-columns: 1fr;
  }

  .cart-summary {
    position: static;
  }
}
</style>
