<script setup lang="ts">
import type { CartItem } from '../stores/cart'

const props = defineProps<{ item: CartItem }>()
const emit = defineEmits<{
  remove: [productId: number]
  update: [productId: number, quantity: number]
}>()

const decrease = () => {
  if (props.item.quantity > 1) {
    emit('update', props.item.product.id, props.item.quantity - 1)
  }
}

const increase = () => {
  emit('update', props.item.product.id, props.item.quantity + 1)
}

const total = () => (props.item.product.price * props.item.quantity).toFixed(2)
</script>

<template>
  <div class="cart-item">
    <div class="item-info">
      <h4 class="item-name">{{ item.product.name }}</h4>
      <span class="item-price">¥{{ item.product.price.toFixed(2) }}</span>
    </div>

    <div class="item-controls">
      <div class="qty-controls">
        <button class="qty-btn" :disabled="item.quantity <= 1" @click="decrease">−</button>
        <span class="qty-value">{{ item.quantity }}</span>
        <button class="qty-btn" @click="increase">+</button>
      </div>

      <span class="item-total">¥{{ total() }}</span>

      <button class="remove-btn" @click="emit('remove', item.product.id)">删除</button>
    </div>
  </div>
</template>

<style scoped>
.cart-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px;
  background: var(--color-card);
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border);
  gap: 16px;
}

.item-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.item-name {
  font-size: 15px;
  font-weight: 600;
  color: var(--color-text);
}

.item-price {
  font-size: 13px;
  color: var(--color-text-secondary);
}

.item-controls {
  display: flex;
  align-items: center;
  gap: 20px;
}

.qty-controls {
  display: flex;
  align-items: center;
  gap: 0;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-sm);
  overflow: hidden;
}

.qty-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  background: var(--color-bg);
  color: var(--color-text);
  font-size: 16px;
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
  width: 40px;
  text-align: center;
  font-size: 14px;
  font-weight: 500;
}

.item-total {
  font-size: 16px;
  font-weight: 700;
  color: var(--color-accent);
  min-width: 80px;
  text-align: right;
}

.remove-btn {
  padding: 6px 12px;
  font-size: 12px;
  color: var(--color-danger);
  background: transparent;
  border: 1px solid var(--color-danger);
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all 0.2s ease;
}

.remove-btn:hover {
  background-color: var(--color-danger);
  color: #fff;
}
</style>
