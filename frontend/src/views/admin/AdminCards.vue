<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '../../api/index'
import { productApi, type Product } from '../../api/product'

interface Card {
  id: number
  content: string
  status: string
  created_at: string
}

const products = ref<Product[]>([])
const selectedProductId = ref<number | null>(null)
const cards = ref<Card[]>([])
const codes = ref('')
const loading = ref(false)
const result = ref('')

const fetchProducts = async () => {
  try {
    const res = await productApi.list()
    products.value = res.data.products || res.data
  } catch {
    products.value = []
  }
}

const fetchCards = async () => {
  if (!selectedProductId.value) return
  try {
    const res = await api.get(`/admin/products/${selectedProductId.value}/cards`)
    cards.value = res.data
  } catch {
    cards.value = []
  }
}

onMounted(fetchProducts)

const importCards = async () => {
  if (!selectedProductId.value || !codes.value.trim()) {
    alert('请选择商品并输入卡密')
    return
  }
  loading.value = true
  result.value = ''
  try {
    const lines = codes.value.split('\n').map(l => l.trim()).filter(Boolean)
    await api.post('/admin/cards/import', {
      product_id: selectedProductId.value,
      contents: lines,
    })
    result.value = `成功导入 ${lines.length} 条卡密`
    codes.value = ''
    fetchCards()
  } catch {
    result.value = '导入失败，请重试'
  } finally {
    loading.value = false
  }
}

const deleteCard = async (id: number) => {
  if (!confirm('确认删除此卡密？')) return
  try {
    await api.delete(`/admin/cards/${id}`)
    fetchCards()
  } catch {
    alert('删除失败')
  }
}

const onProductChange = () => {
  cards.value = []
  fetchCards()
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
        <select v-model="selectedProductId" @change="onProductChange">
          <option :value="null" disabled>请选择商品</option>
          <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }}</option>
        </select>
      </label>

      <label>
        <span>卡密列表（每行一条）</span>
        <textarea v-model="codes" rows="8" placeholder="在此粘贴卡密，每行一条" />
      </label>

      <button class="btn-primary" :disabled="loading" @click="importCards">
        {{ loading ? '导入中...' : '导入卡密' }}
      </button>

      <div v-if="result" class="result-msg">{{ result }}</div>
    </div>

    <div v-if="selectedProductId && cards.length > 0" class="cards-list">
      <h3>已有卡密 ({{ cards.length }})</h3>
      <table class="data-table">
        <thead>
          <tr><th>卡密</th><th>状态</th><th>操作</th></tr>
        </thead>
        <tbody>
          <tr v-for="c in cards" :key="c.id">
            <td class="code">{{ c.content }}</td>
            <td><span class="status-tag" :class="c.status">{{ c.status }}</span></td>
            <td><button class="btn-danger" @click="deleteCard(c.id)">删除</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.admin-cards { max-width: 900px; margin: 0 auto; padding: 40px 24px; }
.page-header { margin-bottom: 32px; }
.page-header h1 { font-size: 24px; font-weight: 600; }
.import-card { background: var(--color-card); border-radius: var(--radius-lg); padding: 32px; box-shadow: 0 1px 3px rgba(0,0,0,0.06); display: flex; flex-direction: column; gap: 20px; margin-bottom: 32px; }
.import-card label { display: flex; flex-direction: column; gap: 6px; }
.import-card label span { font-size: 13px; font-weight: 500; color: var(--color-text-secondary); }
.import-card select, .import-card textarea { padding: 10px 12px; border: 1.5px solid var(--color-border); border-radius: var(--radius-md); font-size: 14px; background: var(--color-bg); }
.import-card textarea { font-family: monospace; resize: vertical; }
.btn-primary { padding: 10px 24px; background: var(--color-primary); color: #fff; border: none; border-radius: var(--radius-full); cursor: pointer; font-size: 14px; align-self: flex-start; }
.btn-primary:hover { background: var(--color-primary-pressed); }
.btn-primary:disabled { opacity: 0.6; cursor: not-allowed; }
.btn-danger { background: none; border: none; color: #e74c3c; cursor: pointer; font-size: 14px; }
.result-msg { padding: 12px 16px; background: var(--color-bg); border-radius: var(--radius-md); font-size: 14px; }
.cards-list h3 { font-size: 16px; font-weight: 600; margin-bottom: 16px; }
.code { font-family: monospace; }
.data-table { width: 100%; border-collapse: collapse; background: var(--color-card); border-radius: var(--radius-lg); overflow: hidden; }
.data-table th, .data-table td { padding: 12px 16px; text-align: left; font-size: 14px; border-bottom: 1px solid var(--color-border); }
.data-table th { background: var(--color-bg); font-weight: 600; color: var(--color-text-secondary); }
.status-tag { padding: 2px 10px; border-radius: var(--radius-full); font-size: 12px; }
.status-tag.unused { background: #d4edda; color: #155724; }
.status-tag.used { background: #f8d7da; color: #721c24; }
</style>
