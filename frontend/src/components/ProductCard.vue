<script setup lang="ts">
import type { Product } from '../api/product'

const props = defineProps<{ product: Product }>()

const categoryLabel: Record<string, string> = {
  game: '游戏',
  vip: 'VIP',
  software: '软件',
  other: '其他',
}

const categoryClass: Record<string, string> = {
  game: 'tag-game',
  vip: 'tag-vip',
  software: 'tag-software',
  other: 'tag-other',
}
</script>

<template>
  <router-link :to="`/product/${product.id}`" class="product-card">
    <div class="card-image">
      <span class="card-icon">✦</span>
    </div>
    <div class="card-body">
      <span class="tag" :class="categoryClass[product.category] || 'tag-other'">
        {{ categoryLabel[product.category] || product.category }}
      </span>
      <h3 class="card-name">{{ product.name }}</h3>
      <p class="card-desc">{{ product.description }}</p>
      <div class="card-footer">
        <span class="card-price">¥{{ product.price.toFixed(2) }}</span>
        <span class="card-stock">库存 {{ product.stock }}</span>
      </div>
    </div>
  </router-link>
</template>

<style scoped>
.product-card {
  display: block;
  background: var(--color-card);
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border);
  overflow: hidden;
  text-decoration: none;
  color: inherit;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}

.product-card:hover {
  transform: translateY(-6px);
  box-shadow: var(--shadow-lg);
}

.card-image {
  height: 160px;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-icon {
  font-size: 48px;
  color: rgba(255, 255, 255, 0.8);
}

.card-body {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card-name {
  font-size: 16px;
  font-weight: 600;
  line-height: 1.4;
  color: var(--color-text);
}

.card-desc {
  font-size: 13px;
  color: var(--color-text-secondary);
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
}

.card-price {
  font-size: 18px;
  font-weight: 700;
  color: var(--color-accent);
}

.card-stock {
  font-size: 12px;
  color: var(--color-text-secondary);
}

.tag-software {
  background-color: rgba(168, 230, 207, 0.2);
  color: #2D7A4F;
}

.tag-other {
  background-color: rgba(255, 139, 148, 0.2);
  color: #C44569;
}
</style>
