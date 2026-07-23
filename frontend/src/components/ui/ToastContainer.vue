<template>
  <Teleport to="body">
    <div class="toast-wrap">
      <TransitionGroup name="toast">
        <div v-for="t in list" :key="t.id" class="toast" :class="`toast--${t.type}`">
          <span class="toast__icon">{{ { success: '✓', error: '✕', info: '♥' }[t.type] }}</span>
          <p class="toast__msg">{{ t.message }}</p>
          <button @click="remove(t.id)" class="toast__close">×</button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup>
import { ref } from 'vue'
const list = ref([])
let nid = 0
function add(message, type = 'info', ms = 3500) {
  const id = ++nid
  list.value.push({ id, message, type })
  setTimeout(() => remove(id), ms)
}
function remove(id) { list.value = list.value.filter(t => t.id !== id) }
defineExpose({ add })
</script>

<style scoped>
.toast-wrap{position:fixed;bottom:1.5rem;right:1.5rem;z-index:9999;display:flex;flex-direction:column;gap:0.75rem;}
.toast{display:flex;align-items:center;gap:0.75rem;padding:1rem 1rem 1rem 1.25rem;background:var(--c-bg-3);border:1px solid var(--c-border);border-radius:var(--r-lg);box-shadow:var(--shadow-elevated);min-width:280px;max-width:380px;}
.toast--success{border-left:3px solid #48c78e;}
.toast--error{border-left:3px solid #ff5b5b;}
.toast--info{border-left:3px solid var(--c-rose);}
.toast__icon{font-size:1rem;flex-shrink:0;}
.toast__msg{flex:1;font-size:0.875rem;}
.toast__close{color:var(--c-text-subtle);font-size:1.2rem;line-height:1;flex-shrink:0;}
.toast__close:hover{color:var(--c-text);}
.toast-enter-active,.toast-leave-active{transition:all 0.3s ease;}
.toast-enter-from{opacity:0;transform:translateX(40px);}
.toast-leave-to{opacity:0;transform:translateX(40px);}
</style>
