<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'
import { productApi, type Product } from '../../api/product'

const products = ref<Product[]>([])
const loading = ref(true)
const showForm = ref(false)
const editingId = ref<number | null>(null)
const form = ref({
  name: '',
  description: '',
  price: 0,
  category: 'game',
  image_url: '',
})

const categories = [
  { value: 'game', label: '游戏' },
  { value: 'vip', label: 'VIP' },
  { value: 'software', label: '软件' },
  { value: 'other', label: '其他' },
]

const fetchProducts = async () => {
  loading.value = true
  try {
    const res = await productApi.list()
    products.value = res.data
  } catch {
    products.value = []
  } finally {
    loading.value = false
  }
}

const openAdd = () => {
  editingId.value = null
  form.value = { name: '', description: '', price: 0, category: 'game', image_url: '' }
  showForm.value = true
}

const openEdit = (p: Product) => {
  editingId.value = p.id
  form.value = {
    name: p.name,
    description: p.description,
    price: p.price,
    category: p.category,
    image_url: p.image_url,
  }
  showForm.value = true
}

const submitForm = async () => {
  try {
    if (editingId.value) {
      await api.put(`/admin/products/${editingId.value}`, form.value)
    } else {
      await api.post('/admin/products', form.value)
    }
    showForm.value = false
    fetchProducts()
  } catch (e: unknown) {
    const msg = e instanceof Error ? e.message : '操作失败'
    alert(msg)
  }
}

const deleteProduct = async (p: Product) => {
  if (!confirm(`确认删除商品「${p.name}」？`)) return
  try {
    await api.delete(`/admin/products/${p.id}`)
    fetchProducts()
  } catch {
    alert('删除失败')
  }
}

onMounted(fetchProducts)
</script>

<template>
  <div class="admin-products">
    <div class="page-header">
      <h1>商品管理</h1>
      <button class="btn-primary" @click="openAdd">添加商品</button>
    </div>

    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="products.length === 0" class="empty-state">暂无商品</div>
    <table v-else class="data-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>名称</th>
          <th>价格</th>
          <th>分类</th>
          <th>库存</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="p in products" :key="p.id">
          <td>{{ p.id }}</td>
          <td>{{ p.name }}</td>
          <td>¥{{ p.price.toFixed(2) }}</td>
          <td>{{ p.category }}</td>
          <td>{{ p.stock }}</td>
          <td class="actions">
            <button class="btn-link" @click="openEdit(p)">编辑</button>
            <button class="btn-danger" @click="deleteProduct(p)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="showForm" class="modal-overlay" @click.self="showForm = false">
      <div class="modal-card">
        <h2>{{ editingId ? '编辑商品' : '添加商品' }}</h2>
        <form @submit.prevent="submitForm" class="admin-form">
          <label>
            <span>名称</span>
            <input v-model="form.name" required />
          </label>
          <label>
            <span>描述</span>
            <textarea v-model="form.description" rows="3" />
          </label>
          <label>
            <span>价格</span>
            <input v-model.number="form.price" type="number" step="0.01" min="0" required />
          </label>
          <label>
            <span>分类</span>
            <select v-model="form.category">
              <option v-for="c in categories" :key="c.value" :value="c.value">{{ c.label }}</option>
            </select>
          </label>
          <label>
            <span>图片URL</span>
            <input v-model="form.image_url" />
          </label>
          <div class="form-actions">
            <button type="button" class="btn-secondary" @click="showForm = false">取消</button>
            <button type="submit" class="btn-primary">保存</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.admin-products {
  max-width: 1100px;
  margin: 0 auto;
  padding: 40px 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.page-header h1 {
  font-size: 24px;
  font-weight: 600;
  color: var(--color-text);
}

.btn-primary {
  padding: 8px 20px;
  background: var(--color-primary);
  color: #fff;
  border: none;
  border-radius: var(--radius-full);
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
}

.btn-primary:hover {
  background: var(--color-primary-pressed);
}

.btn-secondary {
  padding: 8px 20px;
  background: var(--color-card);
  color: var(--color-text-secondary);
  border: 1.5px solid var(--color-border);
  border-radius: var(--radius-full);
  cursor: pointer;
  font-size: 14px;
}

.btn-link {
  background: none;
  border: none;
  color: var(--color-primary);
  cursor: pointer;
  font-size: 14px;
  padding: 4px 8px;
}

.btn-danger {
  background: none;
  border: none;
  color: #e74c3c;
  cursor: pointer;
  font-size: 14px;
  padding: 4px 8px;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  background: var(--color-card);
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.06);
}

.data-table th,
.data-table td {
  padding: 12px 16px;
  text-align: left;
  font-size: 14px;
  border-bottom: 1px solid var(--color-border);
}

.data-table th {
  background: var(--color-bg);
  font-weight: 600;
  color: var(--color-text-secondary);
}

.actions {
  display: flex;
  gap: 4px;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal-card {
  background: var(--color-card);
  border-radius: var(--radius-lg);
  padding: 32px;
  width: 480px;
  max-width: 90vw;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
}

.modal-card h2 {
  margin-bottom: 24px;
  font-size: 20px;
  font-weight: 600;
  color: var(--color-text);
}

.admin-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-form label {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.admin-form label span {
  font-size: 13px;
  font-weight: 500;
  color: var(--color-text-secondary);
}

.admin-form input,
.admin-form select,
.admin-form textarea {
  padding: 10px 12px;
  border: 1.5px solid var(--color-border);
  border-radius: var(--radius-md);
  font-size: 14px;
  background: var(--color-bg);
  color: var(--color-text);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}

.loading-state,
.empty-state {
  text-align: center;
  padding: 60px 0;
  font-size: 16px;
  color: var(--color-text-secondary);
}
</style>
