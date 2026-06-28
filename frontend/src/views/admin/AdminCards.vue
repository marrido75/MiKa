<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'
import { productApi, type Product } from '../../api/product'

const products = ref<Product[]>([])
const productId = ref<number | null>(null)
const codes = ref('')
const loading = ref(false)
const result = ref('')

const fetchProducts = async () => {
  try {
    const res = await productApi.list()
    products.value = res.data
  } catch {
    products.value = []
  }
}

onMounted(fetchProducts)

const importCards = async () => {
  if (!productId.value || !codes.value.trim()) {
    alert('请选择商品并输入卡密')
    return
  }
  loading.value = true
  result.value = ''
  try {
    const lines = codes.value.split('\n').map(l => l.trim()).filter(Boolean)
    const res = await api.post('/admin/cards/import', {
      product_id: productId.value,
      codes: lines,
    })
    result.value = `成功导入 ${res.data.count ?? lines.length} 条卡密`
    codes.value = ''
  } catch {
    result.value = '导入失败，请重试'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="admin-cards">
    <div class="page-header">
      <h1>卡密管理</h1>
    </div>

    <div class="import-card">
      <label>
        <span>选择商品</span>
        <select v-model="productId">
          <option :value="null" disabled>请选择商品</option>
          <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
        </select>
      </label>

      <label>
        <span>卡密列表（每行一条）</span>
        <textarea
          v-model="codes"
          rows="10"
          placeholder="在此粘贴卡密，每行一条"
        />
      </label>

      <button class="btn-primary" :disabled="loading" @click="importCards">
        {{ loading ? '导入中...' : '导入卡密' }}
      </button>

      <div v-if="result" class="result-msg">{{ result }}</div>
    </div>
  </div>
</template>

<style scoped>
.admin-cards {
  max-width: 700px;
  margin: 0 auto;
  padding: 40px 24px;
}

.page-header {
  margin-bottom: 32px;
}

.page-header h1 {
  font-size: 24px;
  font-weight: 600;
  color: var(--color-text);
}

.import-card {
  background: var(--color-card);
  border-radius: var(--radius-lg);
  padding: 32px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.06);
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.import-card label {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.import-card label span {
  font-size: 13px;
  font-weight: 500;
  color: var(--color-text-secondary);
}

.import-card select,
.import-card textarea {
  padding: 10px 12px;
  border: 1.5px solid var(--color-border);
  border-radius: var(--radius-md);
  font-size: 14px;
  background: var(--color-bg);
  color: var(--color-text);
}

.import-card textarea {
  font-family: monospace;
  resize: vertical;
}

.btn-primary {
  padding: 10px 24px;
  background: var(--color-primary);
  color: #fff;
  border: none;
  border-radius: var(--radius-full);
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  align-self: flex-start;
}

.btn-primary:hover {
  background: var(--color-primary-pressed);
}

.btn-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.result-msg {
  padding: 12px 16px;
  background: var(--color-bg);
  border-radius: var(--radius-md);
  font-size: 14px;
  color: var(--color-text);
}
</style>
