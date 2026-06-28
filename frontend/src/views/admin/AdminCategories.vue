<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { categoryApi, type Category } from '../../api/category'

const categories = ref<Category[]>([])
const loading = ref(true)
const showForm = ref(false)
const editingId = ref<number | null>(null)
const form = ref({ name: '', label: '', sort: 0 })

const fetchCategories = async () => {
  loading.value = true
  try {
    const res = await categoryApi.adminList()
    categories.value = res.data
  } catch {
    categories.value = []
  } finally {
    loading.value = false
  }
}

onMounted(fetchCategories)

const openAdd = () => {
  editingId.value = null
  form.value = { name: '', label: '', sort: 0 }
  showForm.value = true
}

const openEdit = (cat: Category) => {
  editingId.value = cat.id
  form.value = { name: cat.name, label: cat.label, sort: cat.sort }
  showForm.value = true
}

const submitForm = async () => {
  try {
    if (editingId.value) {
      await categoryApi.update(editingId.value, form.value)
    } else {
      await categoryApi.create(form.value)
    }
    showForm.value = false
    fetchCategories()
  } catch {
    alert('操作失败')
  }
}

const deleteCategory = async (cat: Category) => {
  if (!confirm(`确认删除分类「${cat.label}」？`)) return
  try {
    await categoryApi.delete(cat.id)
    fetchCategories()
  } catch {
    alert('删除失败')
  }
}
</script>

<template>
  <div class="admin-categories">
    <div class="page-header">
      <h1>分类管理</h1>
      <button class="btn-primary" @click="openAdd">添加分类</button>
    </div>

    <div v-if="loading" class="loading-state">加载中...</div>
    <div v-else-if="categories.length === 0" class="empty-state">暂无分类</div>
    <table v-else class="data-table">
      <thead>
        <tr><th>ID</th><th>标识</th><th>名称</th><th>排序</th><th>状态</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="cat in categories" :key="cat.id">
          <td>{{ cat.id }}</td>
          <td class="code">{{ cat.name }}</td>
          <td>{{ cat.label }}</td>
          <td>{{ cat.sort }}</td>
          <td><span class="status-tag" :class="cat.status">{{ cat.status }}</span></td>
          <td class="actions">
            <button class="btn-link" @click="openEdit(cat)">编辑</button>
            <button class="btn-danger" @click="deleteCategory(cat)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="showForm" class="modal-overlay" @click.self="showForm = false">
      <div class="modal-card">
        <h2>{{ editingId ? '编辑分类' : '添加分类' }}</h2>
        <form @submit.prevent="submitForm" class="admin-form">
          <label>
            <span>标识 (英文)</span>
            <input v-model="form.name" required placeholder="如: game, vip" />
          </label>
          <label>
            <span>名称 (中文)</span>
            <input v-model="form.label" required placeholder="如: 游戏, 会员" />
          </label>
          <label>
            <span>排序 (越小越靠前)</span>
            <input v-model.number="form.sort" type="number" />
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
.admin-categories { max-width: 900px; margin: 0 auto; padding: 40px 24px; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 32px; }
.page-header h1 { font-size: 24px; font-weight: 600; }
.btn-primary { padding: 8px 20px; background: var(--color-primary); color: #fff; border: none; border-radius: var(--radius-full); cursor: pointer; font-size: 14px; }
.btn-primary:hover { background: var(--color-primary-pressed); }
.btn-secondary { padding: 8px 20px; background: var(--color-card); color: var(--color-text-secondary); border: 1.5px solid var(--color-border); border-radius: var(--radius-full); cursor: pointer; font-size: 14px; }
.btn-link { background: none; border: none; color: var(--color-primary); cursor: pointer; font-size: 14px; padding: 4px 8px; }
.btn-danger { background: none; border: none; color: #e74c3c; cursor: pointer; font-size: 14px; padding: 4px 8px; }
.actions { display: flex; gap: 4px; }
.code { font-family: monospace; font-weight: 500; }
.data-table { width: 100%; border-collapse: collapse; background: var(--color-card); border-radius: var(--radius-lg); overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.06); }
.data-table th, .data-table td { padding: 12px 16px; text-align: left; font-size: 14px; border-bottom: 1px solid var(--color-border); }
.data-table th { background: var(--color-bg); font-weight: 600; color: var(--color-text-secondary); }
.status-tag { padding: 2px 10px; border-radius: var(--radius-full); font-size: 12px; }
.status-tag.active { background: #d4edda; color: #155724; }
.status-tag.inactive { background: #f8d7da; color: #721c24; }
.modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.4); display: flex; align-items: center; justify-content: center; z-index: 100; }
.modal-card { background: var(--color-card); border-radius: var(--radius-lg); padding: 32px; width: 420px; max-width: 90vw; }
.modal-card h2 { margin-bottom: 24px; font-size: 20px; font-weight: 600; }
.admin-form { display: flex; flex-direction: column; gap: 16px; }
.admin-form label { display: flex; flex-direction: column; gap: 6px; }
.admin-form label span { font-size: 13px; font-weight: 500; color: var(--color-text-secondary); }
.admin-form input { padding: 10px 12px; border: 1.5px solid var(--color-border); border-radius: var(--radius-md); font-size: 14px; background: var(--color-bg); }
.form-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 8px; }
.loading-state, .empty-state { text-align: center; padding: 60px 0; color: var(--color-text-secondary); }
</style>
