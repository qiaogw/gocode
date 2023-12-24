<template>
  <div @compositionstart="start" @compositionend="end">
    <q-form @submit="onSubmit">
      <q-input
        borderless
        clearable
        outlined
        dense
        debounce="300"
        v-model="filterValue"
        @clear="reset"
        @update:model-value="query"
        placeholder="请输入关键字搜索"
      >
        <template v-slot:append>
          <q-btn icon="search" flat round size="sm" @click="reset">
            <q-tooltip>重置</q-tooltip>
          </q-btn>
        </template>
      </q-input>
    </q-form>
  </div>
</template>
<script setup>
import { ref } from 'vue'

const props = defineProps({
  filter: String,
})
const filterValue = ref('')
const lock = ref(true)
// 锁定中文输入
const emit = defineEmits(['query', 'update:filter'])

const query = (val) => {
  if (lock.value) {
    return
  }
  emit('update:filter', val)
  // emit("query", val);
}
const onSubmit = () => {
  emit('query', { filter: filterValue.value })
}
const reset = () => {
  emit('query', { filter: '' })
}
const start = (val) => {
  lock.value = true
}
const end = (val) => {
  lock.value = false
}
</script>
