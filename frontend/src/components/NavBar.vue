<script setup lang="ts">
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const cart = useCartStore()
const router = useRouter()

const handleLogout = () => {
  auth.logout()
  router.push('/')
}
</script>

<template>
  <nav class="navbar">
    <div class="navbar-inner">
      <router-link to="/" class="navbar-logo">✦ MiKa</router-link>

      <div class="navbar-links">
        <router-link to="/" class="nav-link">首页</router-link>
        <router-link to="/cart" class="nav-link cart-link">
          购物车
          <span v-if="cart.totalCount > 0" class="cart-badge">{{ cart.totalCount }}</span>
        </router-link>
        <router-link v-if="auth.isLoggedIn()" to="/order" class="nav-link">我的订单</router-link>
        <router-link v-if="auth.isAdmin()" to="/admin" class="nav-link admin-link">管理后台</router-link>
      </div>

      <div class="navbar-actions">
        <template v-if="auth.isLoggedIn()">
          <span class="user-name">{{ auth.user?.username }}</span>
          <button class="btn btn-outline btn-sm" @click="handleLogout">退出</button>
        </template>
        <template v-else>
          <router-link to="/login" class="btn btn-outline btn-sm">登录</router-link>
          <router-link to="/register" class="btn btn-primary btn-sm">注册</router-link>
        </template>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: rgba(255, 255, 255, 0.75);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--color-border);
}

.navbar-inner {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.navbar-logo {
  font-size: 20px;
  font-weight: 700;
  color: var(--color-text);
  text-decoration: none;
  font-family: var(--font-serif);
}

.navbar-links {
  display: flex;
  align-items: center;
  gap: 8px;
}

.nav-link {
  padding: 8px 16px;
  font-size: 14px;
  color: var(--color-text-secondary);
  text-decoration: none;
  border-radius: var(--radius-sm);
  transition: all 0.2s ease;
}

.nav-link:hover {
  color: var(--color-text);
  background-color: var(--color-border);
}

.nav-link.router-link-active {
  color: var(--color-primary-pressed);
  font-weight: 500;
}

.cart-link {
  position: relative;
}

.cart-badge {
  position: absolute;
  top: 2px;
  right: 4px;
  min-width: 18px;
  height: 18px;
  padding: 0 5px;
  background-color: var(--color-accent);
  color: #fff;
  font-size: 11px;
  font-weight: 600;
  border-radius: var(--radius-full);
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.navbar-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-name {
  font-size: 14px;
  color: var(--color-text-secondary);
}

.btn-sm {
  padding: 6px 16px;
  font-size: 13px;
}
</style>
