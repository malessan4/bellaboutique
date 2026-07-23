<template>
  <div class="overlay" v-if="modelValue">
    <div class="confirm-card glass">
      <h3 class="confirm-title">{{ title }}</h3>
      <p class="confirm-message">{{ message }}</p>
      <div class="confirm-actions">
        <button class="btn btn-ghost" @click="cancel">{{ cancelText }}</button>
        <button class="btn" :class="confirmClass" @click="confirm" style="background:var(--c-rose);color:white">{{ confirmText }}</button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  modelValue: { type: Boolean, required: true },
  title: { type: String, default: 'Confirmar acción' },
  message: { type: String, required: true },
  confirmText: { type: String, default: 'Aceptar' },
  cancelText: { type: String, default: 'Cancelar' },
  confirmClass: { type: String, default: 'btn-primary' }
})

const emit = defineEmits(['update:modelValue', 'confirm'])

function cancel() {
  emit('update:modelValue', false)
}

function confirm() {
  emit('confirm')
  emit('update:modelValue', false)
}
</script>

<style scoped>
.overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.7); backdrop-filter: blur(4px); -webkit-backdrop-filter: blur(4px); z-index: 1000; display: flex; align-items: center; justify-content: center; animation: fadeIn 0.2s ease; }
.confirm-card { width: 100%; max-width: 400px; padding: var(--sp-8); border-radius: var(--r-xl); text-align: center; }
.confirm-title { font-family: var(--font-display); font-size: 1.5rem; margin-bottom: var(--sp-4); color: var(--c-text); }
.confirm-message { font-size: 0.95rem; color: var(--c-text-muted); margin-bottom: var(--sp-8); line-height: 1.5; }
.confirm-actions { display: flex; justify-content: center; gap: var(--sp-4); }

@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
</style>
