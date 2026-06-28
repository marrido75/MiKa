<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { productApi, type Product } from '../api/product'

const router = useRouter()
const products = ref<Product[]>([])
const loading = ref(true)
const activeCategory = ref('all')
const widgetOpen = ref(false)

const categories = [
  { value: 'all', label: '全部' },
  { value: 'game', label: '游戏' },
  { value: 'vip', label: 'VIP' },
  { value: 'software', label: '软件' },
  { value: 'other', label: '其他' },
]

const stats = [
  { icon: '🛡️', label: '安全可靠' },
  { icon: '⚡', label: '即时发货' },
  { icon: '💬', label: '售后无忧' },
  { icon: '🎯', label: '品质保证' },
]

const fetchProducts = async () => {
  loading.value = true
  try {
    const category = activeCategory.value === 'all' ? undefined : activeCategory.value
    const res = await productApi.list(category)
    products.value = res.data.products || res.data
  } catch {
    products.value = []
  } finally {
    loading.value = false
  }
}

const goDetail = (p: Product) => {
  router.push(`/product/${p.id}`)
}

const stockClass = (stock: number) => {
  if (stock === 0) return 'stock-empty'
  if (stock < 5) return 'stock-critical'
  if (stock < 10) return 'stock-low'
  if (stock < 50) return 'stock-medium'
  return 'stock-high'
}

onMounted(fetchProducts)
</script>

<template>
  <div class="home">
    <!-- 店铺介绍 -->
    <section class="hero-section">
      <div class="hero-inner container">
        <div class="accent-line"></div>
        <h1 class="hero-title">MiKa</h1>
        <p class="hero-slogan">你的数字好物集市</p>
        <p class="hero-desc">精选优质数字商品，安全可靠，即买即用</p>
        <div class="hero-stats">
          <div v-for="(s, i) in stats" :key="i" class="stat-item">
            <span class="stat-icon">{{ s.icon }}</span>
            <span class="stat-label">{{ s.label }}</span>
          </div>
        </div>
      </div>
    </section>

    <!-- 装饰背景 -->
    <div class="hero-bg">
      <div class="circle circle-1"></div>
      <div class="circle circle-2"></div>
      <div class="circle circle-3"></div>
    </div>

    <!-- 商品网格 -->
    <main class="products-section container">
      <div class="section-header">
        <div class="accent-line-sm"></div>
        <h2 class="section-title">精选商品</h2>
        <p class="section-subtitle">为你挑选的优质数字商品</p>
      </div>

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

      <div v-if="loading" class="loading-state">
        <div class="spinner"></div>
        <p>加载中...</p>
      </div>
      <div v-else-if="products.length === 0" class="empty-state">
        <span class="empty-icon">📦</span>
        <p>暂无商品</p>
      </div>
      <div v-else class="product-grid">
        <div
          v-for="p in products"
          :key="p.id"
          class="product-card"
          @click="goDetail(p)"
        >
          <!-- 库存/折扣标签 -->
          <span v-if="p.stock === 0" class="badge badge-soldout">已售罄</span>
          <span v-else-if="p.stock < 10" class="badge badge-lowstock">仅剩 {{ p.stock }}</span>

          <!-- 图片区域 -->
          <div class="card-image">
            <img v-if="p.image_url" :src="p.image_url" :alt="p.name" />
            <div v-else class="image-placeholder">
              <span class="placeholder-icon">✦</span>
            </div>
          </div>

          <!-- 信息区域 -->
          <div class="card-body">
            <span class="card-category tag" :class="`tag-${p.category}`">{{ p.category }}</span>
            <h3 class="card-title">{{ p.name }}</h3>
            <p class="card-desc">{{ p.description }}</p>
            <div class="card-footer">
              <span class="card-price">¥{{ p.price.toFixed(2) }}</span>
              <span class="card-stock" :class="stockClass(p.stock)">库存 {{ p.stock }}</span>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- 浮动客服框 -->
    <div class="float-widget" :class="{ open: widgetOpen }">
      <button class="float-btn" @click="widgetOpen = !widgetOpen">
        <span v-if="!widgetOpen">💬</span>
        <span v-else>✕</span>
      </button>
      <div v-if="widgetOpen" class="float-panel">
        <div class="float-header">
          <h4>联系客服</h4>
          <p>工作时间: 9:00 - 22:00</p>
        </div>
        <div class="float-body">
          <div class="contact-item">
            <span class="contact-icon">📧</span>
            <div>
              <span class="contact-label">邮箱</span>
              <span class="contact-value">support@mika.com</span>
            </div>
          </div>
          <div class="contact-item">
            <span class="contact-icon">💬</span>
            <div>
              <span class="contact-label">微信</span>
              <span class="contact-value">MiKa_Service</span>
            </div>
          </div>
          <div class="contact-item">
            <span class="contact-icon">📱</span>
            <div>
              <span class="contact-label">QQ群</span>
              <span class="contact-value">123456789</span>
            </div>
          </div>
        </div>
        <div class="float-footer">
          <p>常见问题?</p>
          <router-link to="/faq" class="float-link">查看帮助中心 →</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  position: relative;
  overflow: hidden;
}

/* Hero Section */
.hero-section {
  position: relative;
  z-index: 1;
  padding: 80px 0 60px;
}

.hero-inner {
  text-align: center;
}

.accent-line {
  width: 64px;
  height: 3px;
  background: linear-gradient(90deg, var(--color-primary), var(--color-primary-pressed));
  border-radius: 2px;
  margin: 0 auto 24px;
}

.accent-line-sm {
  width: 48px;
  height: 3px;
  background: linear-gradient(90deg, var(--color-primary), var(--color-primary-pressed));
  border-radius: 2px;
  margin: 0 auto 16px;
}

.hero-title {
  font-family: var(--font-serif);
  font-size: 48px;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 12px;
  letter-spacing: -0.5px;
}

.hero-slogan {
  font-size: 20px;
  color: var(--color-primary-pressed);
  font-weight: 500;
  margin-bottom: 12px;
}

.hero-desc {
  font-size: 15px;
  color: var(--color-text-secondary);
  max-width: 480px;
  margin: 0 auto 32px;
  line-height: 1.7;
}

.hero-stats {
  display: flex;
  justify-content: center;
  gap: 40px;
  flex-wrap: wrap;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.stat-icon {
  font-size: 28px;
}

.stat-label {
  font-size: 13px;
  color: var(--color-text-secondary);
  font-weight: 500;
}

/* Decorative circles */
.hero-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 400px;
  pointer-events: none;
  z-index: 0;
}

.circle {
  position: absolute;
  border-radius: 50%;
  opacity: 0.12;
}

.circle-1 {
  width: 300px;
  height: 300px;
  background: var(--color-primary);
  top: -80px;
  right: -60px;
}

.circle-2 {
  width: 200px;
  height: 200px;
  background: var(--color-secondary);
  bottom: -40px;
  left: 10%;
}

.circle-3 {
  width: 120px;
  height: 120px;
  background: var(--color-primary-pressed);
  top: 30%;
  left: 5%;
}

/* Products Section */
.products-section {
  position: relative;
  z-index: 1;
  padding: 40px 24px 80px;
}

.section-header {
  text-align: center;
  margin-bottom: 32px;
}

.section-title {
  font-family: var(--font-serif);
  font-size: 28px;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 8px;
}

.section-subtitle {
  font-size: 14px;
  color: var(--color-text-secondary);
}

/* Category Bar */
.category-bar {
  display: flex;
  gap: 10px;
  margin-bottom: 32px;
  justify-content: center;
  flex-wrap: wrap;
}

.category-btn {
  padding: 10px 24px;
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
  background: var(--color-primary);
  color: #fff;
  border-color: var(--color-primary);
}

/* Product Grid */
.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 16px;
}

.product-card {
  position: relative;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 16px;
  overflow: hidden;
  border: 1px solid var(--color-border);
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.product-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 12px 32px rgba(0, 0, 0, 0.1);
}

/* Badges */
.badge {
  position: absolute;
  top: 12px;
  z-index: 2;
  padding: 3px 10px;
  border-radius: var(--radius-full);
  font-size: 10px;
  font-weight: 600;
  color: #fff;
}

.badge-soldout {
  right: 12px;
  background: #ef4444;
}

.badge-lowstock {
  right: 12px;
  background: #f59e0b;
}

/* Card Image */
.card-image {
  width: 100%;
  height: 140px;
  overflow: hidden;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.image-placeholder {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder-icon {
  font-size: 36px;
  color: rgba(255, 255, 255, 0.7);
}

/* Card Body */
.card-body {
  padding: 14px;
}

.card-category {
  display: inline-block;
  margin-bottom: 8px;
}

.card-title {
  font-size: 14px;
  font-weight: 600;
  color: var(--color-text);
  margin-bottom: 6px;
  line-height: 1.4;
}

.card-desc {
  font-size: 12px;
  color: var(--color-text-secondary);
  line-height: 1.5;
  margin-bottom: 12px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-price {
  font-size: 18px;
  font-weight: 700;
  color: var(--color-accent);
}

.card-stock {
  font-size: 11px;
}

/* Stock colors */
.stock-high { color: #22c55e; }
.stock-medium { color: #3b82f6; }
.stock-low { color: #f59e0b; }
.stock-critical { color: #ef4444; }
.stock-empty { color: #ef4444; font-weight: 600; }

/* States */
.loading-state,
.empty-state {
  text-align: center;
  padding: 80px 0;
  color: var(--color-text-secondary);
}

.empty-icon {
  font-size: 48px;
  display: block;
  margin-bottom: 16px;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--color-border);
  border-top-color: var(--color-primary);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
  margin: 0 auto 16px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Responsive */
@media (max-width: 640px) {
  .hero-title {
    font-size: 36px;
  }
  .hero-stats {
    gap: 24px;
  }
  .product-grid {
    grid-template-columns: 1fr;
  }
}

/* Floating Widget */
.float-widget {
  position: fixed;
  bottom: 24px;
  right: 24px;
  z-index: 1000;
}

.float-btn {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-pressed));
  color: #fff;
  border: none;
  font-size: 24px;
  cursor: pointer;
  box-shadow: 0 4px 20px rgba(168, 230, 207, 0.4);
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.float-btn:hover {
  transform: scale(1.1);
  box-shadow: 0 6px 28px rgba(168, 230, 207, 0.5);
}

.float-panel {
  position: absolute;
  bottom: 72px;
  right: 0;
  width: 300px;
  background: var(--color-card);
  border-radius: var(--radius-lg);
  box-shadow: 0 8px 40px rgba(0, 0, 0, 0.12);
  border: 1px solid var(--color-border);
  overflow: hidden;
  animation: slideUp 0.3s ease;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.float-header {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-pressed));
  padding: 20px;
  color: #fff;
}

.float-header h4 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 4px;
}

.float-header p {
  font-size: 12px;
  opacity: 0.85;
}

.float-body {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  background: var(--color-bg);
  border-radius: var(--radius-md);
}

.contact-icon {
  font-size: 20px;
}

.contact-item > div {
  display: flex;
  flex-direction: column;
}

.contact-label {
  font-size: 11px;
  color: var(--color-text-secondary);
}

.contact-value {
  font-size: 13px;
  font-weight: 500;
  color: var(--color-text);
}

.float-footer {
  padding: 12px 16px;
  border-top: 1px solid var(--color-border);
  text-align: center;
}

.float-footer p {
  font-size: 12px;
  color: var(--color-text-secondary);
  margin-bottom: 6px;
}

.float-link {
  font-size: 13px;
  color: var(--color-primary-pressed);
  text-decoration: none;
  font-weight: 500;
}

.float-link:hover {
  text-decoration: underline;
}
</style>
