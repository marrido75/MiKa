import { createRouter, createWebHistory } from 'vue-router'
import AdminLayout from '../layouts/AdminLayout.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'home', component: () => import('../views/HomeView.vue') },
    { path: '/product/:id', name: 'product', component: () => import('../views/ProductDetailView.vue') },
    { path: '/cart', name: 'cart', component: () => import('../views/CartView.vue') },
    { path: '/order', name: 'orders', component: () => import('../views/OrderView.vue'), meta: { requiresAuth: true } },
    { path: '/login', name: 'login', component: () => import('../views/LoginView.vue') },
    { path: '/register', name: 'register', component: () => import('../views/RegisterView.vue') },
    {
      path: '/admin',
      component: AdminLayout,
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        { path: '', name: 'admin-dashboard', component: () => import('../views/admin/AdminDashboard.vue') },
        { path: 'products', name: 'admin-products', component: () => import('../views/admin/AdminProducts.vue') },
        { path: 'cards', name: 'admin-cards', component: () => import('../views/admin/AdminCards.vue') },
        { path: 'orders', name: 'admin-orders', component: () => import('../views/admin/AdminOrders.vue') },
        { path: 'coupons', name: 'admin-coupons', component: () => import('../views/admin/AdminCoupons.vue') },
        { path: 'users', name: 'admin-users', component: () => import('../views/admin/AdminUsers.vue') },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')

  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else if (to.meta.requiresAdmin && user.role !== 'admin') {
    next('/')
  } else {
    next()
  }
})

export default router
