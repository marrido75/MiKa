<script setup lang="ts">
import { NConfigProvider, NMessageProvider, NDialogProvider } from 'naive-ui'
import { themeOverrides } from './theme'
import NavBar from './components/NavBar.vue'
import FooterBar from './components/FooterBar.vue'
</script>

<template>
  <NConfigProvider :theme-overrides="themeOverrides">
    <NMessageProvider>
      <NDialogProvider>
        <NavBar />
        <main style="min-height: calc(100vh - 140px); padding-top: 64px;">
          <router-view />
        </main>
        <FooterBar />
      </NDialogProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>
