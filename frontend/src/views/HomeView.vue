<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface User {
  id: number
  name: string
  email: string
}

const message = ref('')
const users = ref<User[]>([])
const loading = ref(true)
const error = ref('')

async function fetchData() {
  try {
    loading.value = true
    const [msgRes, usersRes] = await Promise.all([
      fetch('/api/message'),
      fetch('/api/users')
    ])
    const msgData = await msgRes.json()
    const usersData = await usersRes.json()
    message.value = msgData.message
    users.value = usersData
  } catch (e) {
    error.value = 'Failed to fetch data from backend'
    console.error(e)
  } finally {
    loading.value = false
  }
}

onMounted(fetchData)
</script>

<template>
  <div style="max-width: 600px; margin: 40px auto; font-family: sans-serif;">
    <h1>Moka</h1>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error" style="color: red;">{{ error }}</div>

    <div v-else>
      <div style="padding: 20px; background: #f5f5f5; border-radius: 8px; margin-bottom: 20px;">
        <h3>Backend Message</h3>
        <p>{{ message }}</p>
      </div>

      <div style="padding: 20px; background: #f5f5f5; border-radius: 8px;">
        <h3>Users from API</h3>
        <table style="width: 100%; border-collapse: collapse;">
          <thead>
            <tr style="border-bottom: 2px solid #ddd;">
              <th style="text-align: left; padding: 8px;">ID</th>
              <th style="text-align: left; padding: 8px;">Name</th>
              <th style="text-align: left; padding: 8px;">Email</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id" style="border-bottom: 1px solid #ddd;">
              <td style="padding: 8px;">{{ user.id }}</td>
              <td style="padding: 8px;">{{ user.name }}</td>
              <td style="padding: 8px;">{{ user.email }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <button @click="fetchData" style="margin-top: 20px; padding: 8px 16px; cursor: pointer;">
        Refresh  aa a 
      </button>
    </div>
  </div>
</template>
