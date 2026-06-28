<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { productApi, type Product } from '../api/product'
import ProductCard from '../components/ProductCard.vue'

const products = ref<Product[]>([])
const loading = ref(true)
const activeCategory = ref('all')

const categories = [
  { value: 'all', label: '全部' },
  { value: 'game', label: '游戏' },
  { value: 'vip', label: 'VIP' },
  { value: 'software', label: '软件' },
  { value: 'other', label: '其他' },
]

const fetchProducts = async () => {
  loading.value = true
  try {
    const category = activeCategory.value === 'all' ? undefined : activeCategory.value
    const res = await productApi.list(category)
    products.value = res.data
  } catch {
    products.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchProducts)
</script>

<template>
  <div class="home">
    <section class="hero">
      <div class="hero-circle hero-circle-1" />
      <div class="hero-circle hero-circle-2" />
      <div class="hero-circle hero-circle-3" />
      <div class="hero-content">
        <h1 class="hero-title">发现你的数字好物</h1>
        <p class="hero-subtitle">精选优质数字商品，安全可靠，即买即用</p>
      </div>
    </section>

    <section class="products-section">
      <div class="container">
        <div class="category-bar">
          <button
            v-for="cat in categories"
            :key="cat.value"
            class="category-btn"
            :class="{ active: activeCategory === cat.value }"
            @click="activeCategory = cat.value; fetchProducts()"
          >
            {{ cat.label }}
          </button>
        </div>

        <div v-if="loading" class="loading-state">加载中...</div>
        <div v-else-if="products.length === 0" class="empty-state">暂无商品</div>
        <div v-else class="product-grid">
          <ProductCard v-for="p in products" :key="p.id" :product="p" />
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.hero {
  position: relative;
  padding: 100px 24px 60px;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  overflow: hidden;
  text-align: center;
}

.hero-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.15);
}

.hero-circle-1 {
  width: 300px;
  height: 300px;
  top: -80px;
  right: -60px;
}

.hero-circle-2 {
  width: 200px;
  height: 200px;
  bottom: -40px;
  left: 10%;
}

.hero-circle-3 {
  width: 120px;
  height: 120px;
  top: 30%;
  left: 5%;
}

.hero-content {
  position: relative;
  z-index: 1;
}

.hero-title {
  font-family: var(--font-serif);
  font-size: 42px;
  font-weight: 700;
  color: #fff;
  margin-bottom: 12px;
}

.hero-subtitle {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.85);
}

.products-section {
  padding: 40px 0 60px;
}

.category-bar {
  display: flex;
  gap: 8px;
  margin-bottom: 32px;
  flex-wrap: wrap;
}

.category-btn {
  padding: 8px 20px;
  font-size: 14px;
  font-weight: 500;
  border: 1.5px solid var(--color-border);
  border-radius: var(--radius-full);
  background: var(--color-card);
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;
}

.category-btn:hover {
  border-color: var(--color-primary);
  color: var(--color-primary-pressed);
}

.category-btn.active {
  background-color: var(--color-primary);
  color: #fff;
  border-color: var(--color-primary);
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
}

.loading-state,
.empty-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}
</style>
