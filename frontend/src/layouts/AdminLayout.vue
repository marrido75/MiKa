<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()

const menuItems = [
  { path: '/admin', icon: '📊', label: '控制台' },
  { path: '/admin/products', icon: '🛍️', label: '商品管理' },
  { path: '/admin/categories', icon: '🏷️', label: '分类管理' },
  { path: '/admin/cards', icon: '💳', label: '卡密管理' },
  { path: '/admin/orders', icon: '📦', label: '订单管理' },
  { path: '/admin/coupons', icon: '🎫', label: '优惠券' },
  { path: '/admin/users', icon: '👥', label: '用户管理' },
]

const handleLogout = () => {
  auth.logout()
  router.push('/')
}
</script>

<template>
  <div class="admin-layout">
    <aside class="sidebar">
      <div class="sidebar-header">
        <router-link to="/admin" class="logo">✦ MiKa</router-link>
        <span class="admin-badge">管理后台</span>
      </div>

      <nav class="sidebar-nav">
        <router-link
          v-for="item in menuItems"
          :key="item.path"
          :to="item.path"
          class="nav-item"
          :class="{ active: route.path === item.path }"
        >
          <span class="nav-icon">{{ item.icon }}</span>
          <span class="nav-label">{{ item.label }}</span>
        </router-link>
      </nav>

      <div class="sidebar-footer">
        <router-link to="/" class="nav-item">
          <span class="nav-icon">🏠</span>
          <span class="nav-label">返回前台</span>
        </router-link>
        <button class="nav-item logout-btn" @click="handleLogout">
          <span class="nav-icon">🚪</span>
          <span class="nav-label">退出登录</span>
        </button>
      </div>
    </aside>

    <main class="admin-main">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
  background: var(--color-bg);
}

.sidebar {
  width: 240px;
  background: var(--color-card);
  border-right: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  z-index: 100;
}

.sidebar-header {
  padding: 24px 20px;
  border-bottom: 1px solid var(--color-border);
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo {
  font-size: 20px;
  font-weight: 700;
  color: var(--color-text);
  text-decoration: none;
  font-family: var(--font-serif);
}

.admin-badge {
  font-size: 11px;
  padding: 2px 8px;
  background: var(--color-primary);
  color: #fff;
  border-radius: var(--radius-full);
  font-weight: 500;
}

.sidebar-nav {
  flex: 1;
  padding: 16px 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: var(--radius-md);
  color: var(--color-text-secondary);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.15s ease;
  border: none;
  background: none;
  cursor: pointer;
  width: 100%;
  text-align: left;
}

.nav-item:hover {
  background: var(--color-bg);
  color: var(--color-text);
}

.nav-item.active {
  background: var(--color-primary);
  color: #fff;
}

.nav-icon {
  font-size: 18px;
  width: 24px;
  text-align: center;
}

.sidebar-footer {
  padding: 16px 12px;
  border-top: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.logout-btn {
  color: var(--color-danger, #e74c3c);
}

.logout-btn:hover {
  background: rgba(231, 76, 60, 0.1);
}

.admin-main {
  flex: 1;
  margin-left: 240px;
  min-height: 100vh;
}
</style>
