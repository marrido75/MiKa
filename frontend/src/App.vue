<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { NConfigProvider, NMessageProvider, NDialogProvider } from 'naive-ui'
import { themeOverrides } from './theme'
import NavBar from './components/NavBar.vue'
import FooterBar from './components/FooterBar.vue'

const route = useRoute()
const isAdmin = computed(() => route.path.startsWith('/admin'))
</script>

<template>
  <NConfigProvider :theme-overrides="themeOverrides">
    <NMessageProvider>
      <NDialogProvider>
        <template v-if="isAdmin">
          <router-view />
        </template>
        <template v-else>
          <NavBar />
          <main style="min-height: calc(100vh - 140px); padding-top: 64px;">
            <router-view />
          </main>
          <FooterBar />
        </template>
      </NDialogProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>
