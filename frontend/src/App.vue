<template>
  <div id="app-wrapper">
    <template v-if="showLayout">
      <Navbar />
      <main class="main-content">
        <RouterView v-slot="{ Component }">
          <Transition name="page" mode="out-in">
            <component :is="Component" />
          </Transition>
        </RouterView>
      </main>
      <AppFooter />
    </template>
    <template v-else>
      <RouterView />
    </template>
    <ToastContainer ref="toast" />
  </div>
</template>

<script setup>
import { computed, provide, ref } from 'vue'
import { useRoute } from 'vue-router'
import Navbar from '@/components/layout/Navbar.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import ToastContainer from '@/components/ui/ToastContainer.vue'

const route = useRoute()
const toast = ref(null)
const showLayout = computed(() => !(route.path.startsWith('/admin') && route.path !== '/admin/login'))

provide('toast', {
  success: (msg) => toast.value?.add(msg, 'success'),
  error: (msg) => toast.value?.add(msg, 'error'),
  info: (msg) => toast.value?.add(msg, 'info'),
})
</script>

<style>
.main-content { padding-top: 72px; min-height: calc(100vh - 72px); }
</style>
